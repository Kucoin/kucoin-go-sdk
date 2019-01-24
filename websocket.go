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
	signal.Notify(qc, os.Interrupt, syscall.SIGTERM)

	s, err := token.Servers.RandomServer()
	if err != nil {
		return nil, done, ec
	}
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", token.Token)
	u := fmt.Sprintf("%s?%s", s.Endpoint, q.Encode())

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: as.SkipVerifyTls}
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return nil, done, ec
	}

	// Sub-goroutine: stop by external signal
	var stop = false
	go func() {
		<-done
		stop = true
	}()

	// Sub-goroutine: read messages into messages channel
	go func() {
		defer conn.Close()
		defer close(mc)
		defer close(pc)

		for {
			if stop {
				return
			}

			m := &WebSocketDownstreamMessage{}
			if err := conn.ReadJSON(m); err != nil {
				ec <- err
				return
			}
			log.Printf("Received: %s", ToJsonString(m))
			switch m.Type {
			case WelcomeMessage:
			case PongMessage:
				pc <- m.Id
			case AckMessage:
			case ErrorMessage:
				ec <- errors.New(fmt.Sprintf("Message: %s", ToJsonString(m)))
			default:
				mc <- m
			}
		}
	}()

	// Sub-goroutine: keep heartbeat
	go func() {
		// New ticker to send ping message
		pt := time.NewTicker(time.Duration(s.PingInterval) * time.Millisecond)
		defer conn.Close()
		defer pt.Stop()
		defer close(pc)

		for {
			if stop {
				return
			}

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
				case pid := <-pc:
					if pid != p.Id {
						ec <- errors.New(fmt.Sprintf("Invalid pong id %s, expect %s", pid, p.Id))
						return
					}
				case <-time.After(time.Duration(s.PingTimeout) * time.Millisecond):
					ec <- errors.New(fmt.Sprintf("Wait pong timeout in %d ms", s.PingTimeout))
					return
				}
			}
		}
	}()

	// Sub-goroutine: wait to quit signal
	go func() {
		defer close(ec)
		for {
			if stop {
				return
			}
			sg := <-qc
			ec <- errors.New(fmt.Sprintf("Quit due to a signal: %s", sg.String()))
		}
	}()
	return mc, done, ec
}

func (as *ApiService) WebSocketSubscribePublicChannel(topic string, response bool) (<-chan *WebSocketDownstreamMessage, chan<- struct{}, <-chan error) {
	rsp, err := as.WebSocketPublicToken()
	ec := make(chan error)
	done := make(chan struct{})
	if err != nil {
		ec <- err
		return nil, done, ec
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		ec <- err
		return nil, done, ec
	}
	c := NewSubscribeMessage(topic, false, response)
	return as.webSocketSubscribeChannel(t, c)
}

func (as *ApiService) WebSocketSubscribePrivateChannel(topic string, response bool) (<-chan *WebSocketDownstreamMessage, chan<- struct{}, <-chan error) {
	rsp, err := as.WebSocketPrivateToken()
	ec := make(chan error)
	done := make(chan struct{})
	if err != nil {
		ec <- err
		return nil, done, ec
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		ec <- err
		return nil, done, ec
	}
	c := NewSubscribeMessage(topic, true, response)
	return as.webSocketSubscribeChannel(t, c)
}
