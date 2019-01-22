package kucoin

import (
	"math"
	"testing"
	"time"
)

func TestApiService_ServerTime(t *testing.T) {
	s := NewApiServiceFromEnv()
	var ts int64
	_, err := s.ServerTime(&ts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(ts-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
