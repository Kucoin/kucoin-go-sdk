package kucoin

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

// A WebSocketTokenModel contains a token and some servers for WebSocket feed.
type WebSocketTokenModel struct {
	Token   string                `json:"token"`
	Servers WebSocketServersModel `json:"instanceServers"`
}

// A WebSocketServerModel contains some servers for WebSocket feed.
type WebSocketServerModel struct {
	PingInterval int64  `json:"pingInterval"`
	Endpoint     string `json:"endpoint"`
	Protocol     string `json:"protocol"`
	Encrypt      bool   `json:"encrypt"`
	PingTimeout  int64  `json:"pingTimeout"`
}

// A WebSocketServersModel is the set of *WebSocketServerModel.
type WebSocketServersModel []*WebSocketServerModel

// RandomServer returns a server randomly.
func (s WebSocketServersModel) RandomServer() (*WebSocketServerModel, error) {
	l := len(s)
	if l == 0 {
		return nil, errors.New("No available server ")
	}
	return s[rand.Intn(l)], nil
}

// WebSocketPublicToken returns the token for public channel.
func (as *ApiService) WebSocketPublicToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-public", map[string]string{})
	return as.Call(req)
}

// WebSocketPrivateToken returns the token for private channel.
func (as *ApiService) WebSocketPrivateToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-private", map[string]string{})
	return as.Call(req)
}

// All message types of WebSocket.
const (
	WelcomeMessage     = "welcome"
	PingMessage        = "ping"
	PongMessage        = "pong"
	SubscribeMessage   = "subscribe"
	AckMessage         = "ack"
	UnsubscribeMessage = "unsubscribe"
	ErrorMessage       = "error"
)

// A WebSocketMessage represents a message between the WebSocket client and server.
type WebSocketMessage struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// A WebSocketSubscribeMessage represents a message to subscribe the public/private channel.
type WebSocketSubscribeMessage struct {
	*WebSocketMessage
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}

// NewPingMessage creates a ping message instance.
func NewPingMessage() *WebSocketMessage {
	return &WebSocketMessage{
		Id:   IntToString(time.Now().UnixNano()),
		Type: PingMessage,
	}
}

// NewSubscribeMessage creates a subscribe message instance.
func NewSubscribeMessage(topic string, privateChannel, response bool) *WebSocketSubscribeMessage {
	return &WebSocketSubscribeMessage{
		WebSocketMessage: &WebSocketMessage{
			Id:   IntToString(time.Now().UnixNano()),
			Type: SubscribeMessage,
		},
		Topic:          topic,
		PrivateChannel: privateChannel,
		Response:       response,
	}
}

// NewUnsubscribeMessage creates a unsubscribe message instance.
func NewUnsubscribeMessage(topic string, privateChannel, response bool) *WebSocketSubscribeMessage {
	return &WebSocketSubscribeMessage{
		WebSocketMessage: &WebSocketMessage{
			Id:   IntToString(time.Now().UnixNano()),
			Type: UnsubscribeMessage,
		},
		Topic:          topic,
		PrivateChannel: privateChannel,
		Response:       response,
	}
}

// A WebSocketDownstreamMessage represents a message from the WebSocket server to client.
type WebSocketDownstreamMessage struct {
	*WebSocketMessage
	Sn      string          `json:"sn"`
	Topic   string          `json:"topic"`
	Subject string          `json:"subject"`
	RawData json.RawMessage `json:"data"`
}

// ReadData read the data in channel.
func (m *WebSocketDownstreamMessage) ReadData(v interface{}) error {
	return json.Unmarshal(m.RawData, v)
}

// webSocketSubscribeChannel subscribes the specified channel.
func (as *ApiService) webSocketSubscribeChannel(token *WebSocketTokenModel, channel *WebSocketSubscribeMessage) (<-chan *WebSocketDownstreamMessage, chan<- struct{}, <-chan error) {
	var (
		// Stop subscribe channel
		done = make(chan struct{})
		// Quit signal channel
		qc = make(chan os.Signal, 1)
		// Error channel to return
		ec = make(chan error)
		// Pong channel to check pong message
		pc = make(chan string)
		// Downstream message channel
		mc = make(chan *WebSocketDownstreamMessage, 100)
	)
	signal.Notify(qc, os.Interrupt)

	// Find out a server
	s, err := token.Servers.RandomServer()
	if err != nil {
		ec <- err
		return mc, done, ec
	}

	// Concat ws url
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", token.Token)
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	// Ignore verify tls
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: as.apiSkipVerifyTls}

	// Connect ws server
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		ec <- err
		return mc, done, ec
	}

	// Sub-goroutine: read messages into messages channel
	go func() {
		defer conn.Close()
		defer close(mc)
		defer close(pc)

		var subscribeBytes = []byte(ToJsonString(channel))

		for {
			select {
			case <-done:
				return
			default:
				m := &WebSocketDownstreamMessage{}
				if err := conn.ReadJSON(m); err != nil {
					ec <- err
					return
				}
				//log.Printf("ReadJSON: %s", ToJsonString(m))
				switch m.Type {
				case WelcomeMessage:
					if err := conn.WriteMessage(websocket.TextMessage, subscribeBytes); err != nil {
						ec <- err
						return
					}
					//log.Printf("Subscribing: %s, %s", channel.Id, channel.Topic)
				case PongMessage:
					pc <- m.Id
				case AckMessage:
					//log.Printf("Subscribed: %s==%s? %s", channel.Id, m.Id, channel.Topic)
				case ErrorMessage:
					ec <- errors.Errorf("Error message: %s", ToJsonString(m))
					return
				default:
					mc <- m
				}
			}
		}
	}()

	// Sub-goroutine: keep heartbeat
	go func() {
		// New ticker to send ping message
		pt := time.NewTicker(time.Duration(s.PingInterval)*time.Millisecond - time.Second)
		defer pt.Stop()

		for {
			select {
			case <-done:
				return
			case <-pt.C:
				p := NewPingMessage()
				m := ToJsonString(p)
				if err := conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
					ec <- err
					return
				}

				//log.Printf("Ping: %s", ToJsonString(p))
				// Waiting (with timeout) for the server to response pong message
				// If timeout, close this connection
				select {
				case pid := <-pc:
					if pid != p.Id {
						ec <- errors.Errorf("Invalid pong id %s, expect %s", pid, p.Id)
						return
					}
				case <-time.After(time.Duration(s.PingTimeout) * time.Millisecond):
					ec <- errors.Errorf("Wait pong message timeout in %d ms", s.PingTimeout)
					return
				}
			}
		}
	}()

	// Sub-goroutine: wait to quit signal
	go func() {
		defer close(ec)
		select {
		case <-done:
		case sg := <-qc:
			ec <- errors.Errorf("Quit due to a signal: %s", sg.String())
		}
	}()
	return mc, done, ec
}

// WebSocketSubscribePublicChannel subscribes the specified public channel.
func (as *ApiService) WebSocketSubscribePublicChannel(topic string, response bool) (<-chan *WebSocketDownstreamMessage, chan<- struct{}, <-chan error) {
	rsp, err := as.WebSocketPublicToken()
	mc := make(<-chan *WebSocketDownstreamMessage)
	ec := make(chan error)
	done := make(chan<- struct{})
	if err != nil {
		ec <- err
		return mc, done, ec
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		ec <- err
		return mc, done, ec
	}
	m := NewSubscribeMessage(topic, false, response)
	return as.webSocketSubscribeChannel(t, m)
}

// WebSocketSubscribePrivateChannel subscribes the specified private channel.
func (as *ApiService) WebSocketSubscribePrivateChannel(topic string, response bool) (<-chan *WebSocketDownstreamMessage, chan<- struct{}, <-chan error) {
	rsp, err := as.WebSocketPrivateToken()
	mc := make(<-chan *WebSocketDownstreamMessage)
	ec := make(chan error)
	done := make(chan<- struct{})
	if err != nil {
		ec <- err
		return mc, done, ec
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		ec <- err
		return mc, done, ec
	}
	m := NewSubscribeMessage(topic, true, response)
	return as.webSocketSubscribeChannel(t, m)
}
