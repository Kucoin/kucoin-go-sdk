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
	mc, done, ec := s.WebSocketSubscribePublicChannel("/market/ticker:BTC-USDT", false)
	defer close(done) // Stop subscribe
	var i = 0
	for {
		select {
		case m := <-mc:
			t.Log(ToJsonString(m))
			i++
			if i == 100 {
				return
			}
		case err := <-ec:
			t.Fatal(err)
		}
	}
}

func TestApiService_WebSocketSubscribePrivateChannel(t *testing.T) {
	t.SkipNow()
	s := NewApiServiceFromEnv()
	s.SkipVerifyTls = true
	mc, done, ec := s.WebSocketSubscribePublicChannel("/market/ticker:BTC-USDT", false)
	defer close(done) // Stop subscribe
	var i = 0
	for {
		select {
		case m := <-mc:
			t.Log(ToJsonString(m))
			i++
			if i == 100 {
				return
			}
		case err := <-ec:
			t.Fatal(err)
		}
	}
}
