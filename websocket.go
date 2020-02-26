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
	"github.com/sirupsen/logrus"
)

// A WebSocketTokenModel contains a token and some servers for WebSocket feed.
type WebSocketTokenModel struct {
	Token             string                `json:"token"`
	Servers           WebSocketServersModel `json:"instanceServers"`
	AcceptUserMessage bool                  `json:"accept_user_message"`
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
	Message            = "message"
	Notice             = "notice"
	Command            = "command"
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
func NewSubscribeMessage(topic string, privateChannel bool) *WebSocketSubscribeMessage {
	return &WebSocketSubscribeMessage{
		WebSocketMessage: &WebSocketMessage{
			Id:   IntToString(time.Now().UnixNano()),
			Type: SubscribeMessage,
		},
		Topic:          topic,
		PrivateChannel: privateChannel,
		Response:       true,
	}
}

// A WebSocketUnsubscribeMessage represents a message to unsubscribe the public/private channel.
type WebSocketUnsubscribeMessage WebSocketSubscribeMessage

// NewUnsubscribeMessage creates a unsubscribe message instance.
func NewUnsubscribeMessage(topic string, privateChannel bool) *WebSocketUnsubscribeMessage {
	return &WebSocketUnsubscribeMessage{
		WebSocketMessage: &WebSocketMessage{
			Id:   IntToString(time.Now().UnixNano()),
			Type: UnsubscribeMessage,
		},
		Topic:          topic,
		PrivateChannel: privateChannel,
		Response:       true,
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
	// ACK channel to check pong message
	acks chan string
	// Error channel
	errors chan error
	// Downstream message channel
	messages        chan *WebSocketDownstreamMessage
	conn            *websocket.Conn
	token           *WebSocketTokenModel
	server          *WebSocketServerModel
	enableHeartbeat bool
	skipVerifyTls   bool
	timeout         time.Duration
}

var defaultTimeout = time.Second * 5

// WebSocketClientOpts defines the options for the client
// during the websocket connection.
type WebSocketClientOpts struct {
	Token         *WebSocketTokenModel
	TLSSkipVerify bool
	Timeout       time.Duration
}

// NewWebSocketClient creates an instance of WebSocketClient.
func (as *ApiService) NewWebSocketClient(token *WebSocketTokenModel) *WebSocketClient {
	return as.NewWebSocketClientOpts(WebSocketClientOpts{
		Token:         token,
		TLSSkipVerify: as.apiSkipVerifyTls,
		Timeout:       defaultTimeout,
	})
}

// NewWebSocketClientOpts creates an instance of WebSocketClient with the parsed options.
func (as *ApiService) NewWebSocketClientOpts(opts WebSocketClientOpts) *WebSocketClient {
	wc := &WebSocketClient{
		wg:            &sync.WaitGroup{},
		done:          make(chan struct{}),
		errors:        make(chan error, 1),
		pongs:         make(chan string, 1),
		acks:          make(chan string, 1),
		token:         opts.Token,
		messages:      make(chan *WebSocketDownstreamMessage, 2048),
		skipVerifyTls: opts.TLSSkipVerify,
		timeout:       opts.Timeout,
	}
	return wc
}

// Connect connects the WebSocket server.
func (wc *WebSocketClient) Connect() (<-chan *WebSocketDownstreamMessage, <-chan error, error) {
	// Find out a server
	s, err := wc.token.Servers.RandomServer()
	if err != nil {
		return wc.messages, wc.errors, err
	}
	wc.server = s

	// Concat ws url
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", wc.token.Token)
	if wc.token.AcceptUserMessage == true {
		q.Add("acceptUserMessage", "true")
	}
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	// Ignore verify tls
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: wc.skipVerifyTls}

	// Connect ws server
	websocket.DefaultDialer.ReadBufferSize = 2048000 //2000 kb
	wc.conn, _, err = websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return wc.messages, wc.errors, err
	}

	// Must read the first welcome message
	for {
		m := &WebSocketDownstreamMessage{}
		if err := wc.conn.ReadJSON(m); err != nil {
			return wc.messages, wc.errors, err
		}
		if DebugMode {
			logrus.Debugf("Received a WebSocket message: %s", ToJsonString(m))
		}
		if m.Type == ErrorMessage {
			return wc.messages, wc.errors, errors.Errorf("Error message: %s", ToJsonString(m))
		}
		if m.Type == WelcomeMessage {
			break
		}
	}

	wc.wg.Add(2)
	go wc.read()
	go wc.keepHeartbeat()

	return wc.messages, wc.errors, nil
}

func (wc *WebSocketClient) read() {
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
			if DebugMode {
				logrus.Debugf("Received a WebSocket message: %s", ToJsonString(m))
			}
			// log.Printf("ReadJSON: %s", ToJsonString(m))
			switch m.Type {
			case WelcomeMessage:
			case PongMessage:
				if wc.enableHeartbeat {
					wc.pongs <- m.Id
				}
			case AckMessage:
				// log.Printf("Subscribed: %s==%s? %s", channel.Id, m.Id, channel.Topic)
				wc.acks <- m.Id
			case ErrorMessage:
				wc.errors <- errors.Errorf("Error message: %s", ToJsonString(m))
				return
			case Message, Notice, Command:
				wc.messages <- m
			default:
				wc.errors <- errors.Errorf("Unknown message type: %s", m.Type)
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
			if DebugMode {
				logrus.Debugf("Sent a WebSocket message: %s", m)
			}
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
func (wc *WebSocketClient) Subscribe(channels ...*WebSocketSubscribeMessage) error {
	for _, c := range channels {
		m := ToJsonString(c)
		if DebugMode {
			logrus.Debugf("Sent a WebSocket message: %s", m)
		}
		if err := wc.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			return err
		}
		//log.Printf("Subscribing: %s, %s", c.Id, c.Topic)
		select {
		case id := <-wc.acks:
			//log.Printf("ack: %s=>%s", id, c.Id)
			if id != c.Id {
				return errors.Errorf("Invalid ack id %s, expect %s", id, c.Id)
			}
		case err := <-wc.errors:
			return errors.Errorf("Subscribe failed, %s", err.Error())
		case <-time.After(wc.timeout):
			return errors.Errorf("Wait ack message timeout in %v", wc.timeout)
		}
	}
	return nil
}

// Unsubscribe unsubscribes the specified channel.
func (wc *WebSocketClient) Unsubscribe(channels ...*WebSocketUnsubscribeMessage) error {
	for _, c := range channels {
		m := ToJsonString(c)
		if DebugMode {
			logrus.Debugf("Sent a WebSocket message: %s", m)
		}
		if err := wc.conn.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			return err
		}
		//log.Printf("Unsubscribing: %s, %s", c.Id, c.Topic)
		select {
		case id := <-wc.acks:
			//log.Printf("ack: %s=>%s", id, c.Id)
			if id != c.Id {
				return errors.Errorf("Invalid ack id %s, expect %s", id, c.Id)
			}
		case <-time.After(wc.timeout):
			return errors.Errorf("Wait ack message timeout in %v", wc.timeout)
		}
	}
	return nil
}

// Stop stops subscribing the specified channel, all goroutines quit.
func (wc *WebSocketClient) Stop() {
	close(wc.done)
	_ = wc.conn.Close()
	wc.wg.Wait()
}
