package kucoin

import (
	"math"
	"testing"
	"time"
)

func TestApiService_ServerTime(t *testing.T) {
	s := NewApiServiceFromEnv()
	ts, err := s.ServerTime()
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(ts-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
