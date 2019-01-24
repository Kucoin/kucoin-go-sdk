package kucoin

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
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

func (as *ApiService) webSocketSubscribeChannel(token *WebSocketTokenModel, channel *WebSocketSubscribeMessage) error {
	server, err := token.Servers.RandomServer()
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Add("connectId", IntToString(time.Now().UnixNano()))
	q.Add("token", token.Token)
	u := fmt.Sprintf("%s?%s", server.Endpoint, q.Encode())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: as.SkipVerifyTls}
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return err
	}
	defer c.Close()

	// Done channel to close sub-goroutine
	done := make(chan error)
	// Pong channel to check pong message
	pc := make(chan string)

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Read: %s", err.Error())
				done <- err
				return
			}
			log.Printf("Recv: %s", message)
		}
	}()

	// New ticker to send ping message
	pt := time.NewTicker(time.Duration(server.PingInterval) * time.Millisecond)
	defer pt.Stop()

	for {
		select {
		case err := <-done:
			return err
		case <-pt.C:
			p := NewPingMessage()
			err := c.WriteMessage(websocket.TextMessage, []byte(ToJsonString(p)))
			if err != nil {
				return err
			}
			// Waiting (with timeout) for the server to response pong message
			// If timeout, close this connection
			select {
			case pid := <-pc:
				if pid != p.Id {
					return errors.New(fmt.Sprintf("Invalid pong id %s, expect %s", pid, p.Id))
				}
			case <-time.After(time.Duration(server.PingTimeout) * time.Millisecond):
				return errors.New(fmt.Sprintf("Wait pong timeout in %d ms", server.PingTimeout))
			}
		case <-interrupt:
			if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return errors.New("Interrupted ")
		}
	}
}

func (as *ApiService) WebSocketSubscribePublicChannel(topic string, response bool) error {
	rsp, err := as.WebSocketPublicToken()
	if err != nil {
		return nil
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		return err
	}
	c := NewSubscribeMessage(topic, false, response)
	return as.webSocketSubscribeChannel(t, c)
}

func (as *ApiService) WebSocketSubscribePrivateChannel(topic string, response bool) error {
	rsp, err := as.WebSocketPrivateToken()
	if err != nil {
		return nil
	}
	t := &WebSocketTokenModel{}
	if err := rsp.ReadData(t); err != nil {
		return err
	}
	c := NewSubscribeMessage(topic, true, response)
	return as.webSocketSubscribeChannel(t, c)
}
