package kucoin

import (
	"testing"
)

func TestApiService_Candles(t *testing.T) {
	s := NewApiServiceFromEnv()
	params := map[string]string{
		"startEnd": "1649472662",
		"startAt":  "1649472895",
		"type":     "12hour",
		"symbol":   "BTC-USDT",
	}
	rsp, err := s.CandleHistory(params)
	if err != nil {
		t.Fatal(err)
	}
	cm := make([]*Candle, 0, 10)
	if err := rsp.ReadData(&cm); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(cm))

}
