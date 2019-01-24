package kucoin

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketTokenModel struct {
	Token   string                `json:"token"`
	Servers WebSocketServersModel `json:"instanceServers"`
}
type WebSocketServerModel struct {
	PingInterval int64  `json:"pingInterval"`
	Endpoint     string `json:"endpoint"`
	Protocol     string `json:"protocol"`
	Encrypt      bool   `json:"encrypt"`
	PingTimeout  int64  `json:"pingTimeout"`
}

type WebSocketServersModel []*WebSocketServerModel

func (s WebSocketServersModel) RandomServer() (*WebSocketServerModel, error) {
	l := len(s)
	if l == 0 {
		return nil, errors.New("No available server ")
	}
	return s[rand.Intn(l)], nil
}

func (as *ApiService) WebSocketPublicToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-public", map[string]string{})
	return as.call(req)
}

func (as *ApiService) WebSocketPrivateToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-private", map[string]string{})
	return as.call(req)
}

const (
	WelcomeMessage     = "welcome"
	PingMessage        = "ping"
	PongMessage        = "pong"
	SubscribeMessage   = "subscribe"
	AckMessage         = "ack"
	UnsubscribeMessage = "unsubscribe"
	ErrorMessage       = "error"
)

type WebSocketMessage struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type WebSocketSubscribeMessage struct {
	*WebSocketMessage
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}

func NewPingMessage() *WebSocketMessage {
	return &WebSocketMessage{
		Id:   IntToString(time.Now().UnixNano()),
		Type: PingMessage,
	}
}

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

type WebSocketDownstreamMessage struct {
	*WebSocketMessage
	Topic   string          `json:"topic"`
	Subject bool            `json:"subject"`
	RawData json.RawMessage `json:"data"`
}

func (m *WebSocketDownstreamMessage) ReadData(v interface{}) error {
	if err := json.Unmarshal(m.RawData, v); err != nil {
		return err
	}
	return nil
}

func (as *ApiService) webSocketSubscribeChannel(token *WebSocketTokenModel, channel *WebSocketSubscribeMessage) (chan *WebSocketDownstreamMessage, error) {
	s, err := token.Servers.RandomServer()
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", token.Token)
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: as.SkipVerifyTls}
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return nil, err
	}

	var (
		// Quit signals
		quit = make(chan os.Signal, 1)
		// Error channel to return
		ec = make(chan error)
		// Pong channel to check pong message
		pong = make(chan string)
		// Downstream message channel
		messages = make(chan *WebSocketDownstreamMessage, 100)
	)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	// Sub-goroutine: read messages into messages channel
	go func() {
		defer conn.Close()

		for {
			m := &WebSocketDownstreamMessage{}
			if err := conn.ReadJSON(m); err != nil {
				ec <- err
				return
			}
			log.Printf("Received: %s", ToJsonString(m))
			switch m.Type {
			case WelcomeMessage:
			case PongMessage:
				pong <- m.Id
			case AckMessage:
			case ErrorMessage:
				ec <- errors.New(fmt.Sprintf("Message: %s", ToJsonString(m)))
			default:
				messages <- m
			}
		}
	}()

	// Sub-goroutine: Keep heartbeat & handle quit signal
	go func() {
		// New ticker to send ping message
		pt := time.NewTicker(time.Duration(s.PingInterval) * time.Millisecond)
		defer pt.Stop()
		defer conn.Close()

		for {
			select {
			case <-pt.C:
				p := NewPingMessage()
				log.Println("Send ping: ", p.Id)
				if err := conn.WriteJSON(p); err != nil {
					ec <- err
					return
				}
				// Waiting (with timeout) for the server to response pong message
				// If timeout, close this connection
				select {
				case pid := <-pong:
					if pid != p.Id {
						ec <- errors.New(fmt.Sprintf("Invalid pong id %s, expect %s", pid, p.Id))
						return
					}
				case <-time.After(time.Duration(s.PingTimeout) * time.Millisecond):
					ec <- errors.New(fmt.Sprintf("Wait pong timeout in %d ms", s.PingTimeout))
					return
				}
			case s := <-quit:
				ec <- errors.New(fmt.Sprintf("Quit by signal: %s", s.String()))
				return
			}
		}
	}()
	return messages, err
}

func (as *ApiService) WebSocketSubscribePublicChannel(topic string, response bool) (chan *WebSocketDownstreamMessage, error) {
	rsp, err := as.WebSocketPublicToken()
	if err != nil {
		return nil, err
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		return nil, err
	}
	c := NewSubscribeMessage(topic, false, response)
	return as.webSocketSubscribeChannel(t, c)
}

func (as *ApiService) WebSocketSubscribePrivateChannel(topic string, response bool) (chan *WebSocketDownstreamMessage, error) {
	rsp, err := as.WebSocketPrivateToken()
	if err != nil {
		return nil, err
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		return nil, err
	}
	c := NewSubscribeMessage(topic, true, response)
	return as.webSocketSubscribeChannel(t, c)
}
