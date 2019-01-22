package kucoin

import (
	"math"
	"testing"
	"time"
)

func TestApiService_ServerTime(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.ServerTime()
	if err != nil {
		t.Fatal(err)
	}
	var ts int64
	rsp.ReadData(&ts)
	t.Log(ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(ts-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
