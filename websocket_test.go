package kucoin

import (
	"testing"
)

func TestApiService_WebSocketPublicToken(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		t.Fatal(err)
	}
	pt := &WebSocketTokenModel{}
	if err := rsp.ReadData(pt); err != nil {
		t.Fatal(err)
	}
	t.Log(pt.Token)
	switch {
	case pt.Token == "":
		t.Error("Empty key 'token'")
	case len(pt.Servers) == 0:
		t.Fatal("Empty key 'instanceServers'")
	}
	for _, s := range pt.Servers {
		t.Log(ToJsonString(s))
		switch {
		case s.Endpoint == "":
			t.Error("Empty key 'endpoint'")
		case s.Protocol == "":
			t.Fatal("Empty key 'protocol'")
		}
	}
}

func TestApiService_WebSocketPrivateToken(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.WebSocketPrivateToken()
	if err != nil {
		t.Fatal(err)
	}
	pt := &WebSocketTokenModel{}
	if err := rsp.ReadData(pt); err != nil {
		t.Fatal(err)
	}
	t.Log(pt.Token)
	switch {
	case pt.Token == "":
		t.Error("Empty key 'token'")
	case len(pt.Servers) == 0:
		t.Fatal("Empty key 'instanceServers'")
	}
	for _, s := range pt.Servers {
		t.Log(ToJsonString(s))
		switch {
		case s.Endpoint == "":
			t.Error("Empty key 'endpoint'")
		case s.Protocol == "":
			t.Fatal("Empty key 'protocol'")
		}
	}
}

func TestWebSocketClient_Connect(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		t.Fatal(err)
	}

	tk := &WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}

	ch := NewSubscribeMessage("/market/ticker:KCS-BTC", false, true)

	c := s.NewWebSocketClient(tk, ch)

	if err := c.Connect(); err != nil {
		t.Fatal(err)
	}
}
func TestWebSocketClient_Subscribe(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()

	rsp, err := s.WebSocketPublicToken()
	if err != nil {
		t.Fatal(err)
	}

	tk := &WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		t.Fatal(err)
	}

	ch1 := NewSubscribeMessage("/market/ticker:KCS-BTC", false, true)
	ch2 := NewSubscribeMessage("/market/ticker:ETH-BTC", false, true)

	c := s.NewWebSocketClient(tk, ch1, ch2)

	if err := c.Connect(); err != nil {
		t.Fatal(err)
	}

	mc, ec := c.Subscribe()

	var i = 0
	for {
		select {
		case err := <-ec:
			c.Stop() // Stop subscribing the WebSocket feed
			t.Fatal(err)
		case msg := <-mc:
			t.Log(ToJsonString(msg))
			i++
			if i == 10 {
				t.Log("Exit test")
				c.Stop() // Stop subscribing the WebSocket feed
				return
			}
		}
	}
}
