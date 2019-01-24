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

func TestApiService_WebSocketSubscribePublicChannel(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	s.SkipVerifyTls = true
	mc, err := s.WebSocketSubscribePublicChannel("/market/ticker:BTC-USDT", false)
	if err != nil {
		t.Error(err)
	}
	for {
		m := <-mc
		t.Log(ToJsonString(m))
	}
}

func TestApiService_WebSocketSubscribePrivateChannel(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	s.SkipVerifyTls = true
	mc, err := s.WebSocketSubscribePublicChannel("/market/ticker:BTC-USDT", false)
	if err != nil {
		t.Error(err)
	}
	for {
		m := <-mc
		t.Log(ToJsonString(m))
	}
}
