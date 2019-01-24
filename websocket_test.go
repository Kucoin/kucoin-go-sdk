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

func TestApiService_WebSocketSubscribeChannel(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	s.SkipVerifyTls = true
	c := NewSubscribeMessage("/market/snapshot:BTC-USDT", false, false)
	if err := s.WebSocketSubscribePublicChannel(c); err != nil {
		t.Error(err)
	}
}
