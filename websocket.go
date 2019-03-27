package kucoin

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
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

// A WebSocketClient represents a connection to WebSocket server.
type WebSocketClient struct {
	// Wait all goroutines quit
	wg *sync.WaitGroup
	// Stop subscribing channel
	done chan struct{}
	// Pong channel to check pong message
	pongs chan string
	// Error channel
	errors chan error
	// Downstream message channel
	messages        chan *WebSocketDownstreamMessage
	conn            *websocket.Conn
	token           *WebSocketTokenModel
	channels        []*WebSocketSubscribeMessage
	server          *WebSocketServerModel
	enableHeartbeat bool
	skipVerifyTls   bool
}

// NewWebSocketClient creates an instance of WebSocketClient.
func (as *ApiService) NewWebSocketClient(token *WebSocketTokenModel, channel ...*WebSocketSubscribeMessage) *WebSocketClient {
	wc := &WebSocketClient{
		wg:            &sync.WaitGroup{},
		done:          make(chan struct{}),
		errors:        make(chan error, 1),
		pongs:         make(chan string, 1),
		channels:      channel,
		token:         token,
		messages:      make(chan *WebSocketDownstreamMessage, 100),
		skipVerifyTls: as.apiSkipVerifyTls,
	}
	return wc
}

// Connect connects the WebSocket server.
func (wc *WebSocketClient) Connect() error {
	// Find out a server
	s, err := wc.token.Servers.RandomServer()
	if err != nil {
		return err
	}
	wc.server = s

	// Concat ws url
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", wc.token.Token)
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	// Ignore verify tls
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: wc.skipVerifyTls}

	// Connect ws server
	wc.conn, _, err = websocket.DefaultDialer.Dial(u, nil)
	return err
}

func (wc *WebSocketClient) subscribe() {
	defer func() {
		close(wc.pongs)
		close(wc.messages)
		wc.wg.Done()
	}()

	for {
		select {
		case <-wc.done:
			return
		default:
			m := &WebSocketDownstreamMessage{}
			if err := wc.conn.ReadJSON(m); err != nil {
				wc.errors <- err
				return
			}
			// log.Printf("ReadJSON: %s", ToJsonString(m))
			switch m.Type {
			case WelcomeMessage:
				for _, c := range wc.channels {
					if err := wc.conn.WriteMessage(websocket.TextMessage, []byte(ToJsonString(c))); err != nil {
						wc.errors <- err
						return
					}
					// log.Printf("Subscribing: %s, %s", c.Id, c.Topic)
				}
			case PongMessage:
				if wc.enableHeartbeat {
					wc.pongs <- m.Id
				}
			case AckMessage:
				// log.Printf("Subscribed: %s==%s? %s", channel.Id, m.Id, channel.Topic)
			case ErrorMessage:
				wc.errors <- errors.Errorf("Error message: %s", ToJsonString(m))
				return
			default:
				wc.messages <- m
			}
		}
	}
}

func (wc *WebSocketClient) keepHeartbeat() {
	wc.enableHeartbeat = true
	// New ticker to send ping message
	pt := time.NewTicker(time.Duration(wc.server.PingInterval)*time.Millisecond - time.Millisecond*200)
	defer wc.wg.Done()
	defer pt.Stop()

	for {
		select {
		case <-wc.done:
			return
		case <-pt.C:
			p := NewPingMessage()
			m := ToJsonString(p)
			if err := wc.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
				wc.errors <- err
				return
			}

			// log.Printf("Ping: %s", ToJsonString(p))
			// Waiting (with timeout) for the server to response pong message
			// If timeout, close this connection
			select {
			case pid := <-wc.pongs:
				if pid != p.Id {
					wc.errors <- errors.Errorf("Invalid pong id %s, expect %s", pid, p.Id)
					return
				}
			case <-time.After(time.Duration(wc.server.PingTimeout) * time.Millisecond):
				wc.errors <- errors.Errorf("Wait pong message timeout in %d ms", wc.server.PingTimeout)
				return
			}
		}
	}
}

// Subscribe subscribes the specified channel.
func (wc *WebSocketClient) Subscribe() (<-chan *WebSocketDownstreamMessage, <-chan error) {
	wc.wg.Add(2)
	go wc.subscribe()
	go wc.keepHeartbeat()
	return wc.messages, wc.errors
}

// Stop stops subscribing the specified channel, all goroutines quit.
func (wc *WebSocketClient) Stop() {
	close(wc.done)
	_ = wc.conn.Close()
	wc.wg.Wait()
}
