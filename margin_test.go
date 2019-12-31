package kucoin

import (
	"math/rand"
	"testing"
	"time"
)

func TestApiService_CurrentMarkPrice(t *testing.T) {

	symbols := []string{"USDT-BTC", "ETH-BTC", "LTC-BTC", "EOS-BTC", "XRP-BTC", "KCS-BTC"}
	rand.Seed(time.Now().UnixNano())
	symbol := symbols[rand.Intn(len(symbols))]

	s := NewApiServiceFromEnv()
	rsp, err := s.CurrentMarkPrice(symbol)

	if err != nil {
		t.Fatal(err)
	}
	o := &MarkPriceModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Symbol == "":
		t.Error("empty key 'Symbol'")
	}
}

func TestApiService_MarginConfig(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.MarginConfig()

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginConfigModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))
	switch {
	case o.MaxLeverage == "":
		t.Error("empty key 'MaxLeverage'")
	}
}

func TestApiService_MarginAccount(t *testing.T) {

	s := NewApiServiceFromEnv()
	rsp, err := s.MarginAccount()

	if err != nil {
		t.Fatal(err)
	}

	o := &MarginAccountModel{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}

	t.Log(ToJsonString(o))

	switch {
	case len(o.Accounts) == 0:
		t.Error("empty key 'Accounts'")
	}
}
