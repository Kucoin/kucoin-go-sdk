package kucoin

import (
	"math"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	ts, err := Timestamp()
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(ts-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
