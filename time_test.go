package kucoin

import (
	"context"
	"math"
	"testing"
	"time"
)

func TestApiService_ServerTime(t *testing.T) {
	s := NewApiServiceFromEnv()
	rsp, err := s.ServerTime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var ts ServerTimeModel
	if err := rsp.ReadData(&ts); err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(int64(ts)-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
