package kucoin

import (
	"log"
	"math"
	"testing"
	"time"
)

func TestTime_Timestamp(t *testing.T) {
	tm := Time{}
	ts, err := tm.Timestamp()
	if err != nil {
		t.Error(err.Error())
	}
	log.Println(time.Now().UnixNano()/1000/1000, ts)
	now := time.Now().UnixNano() / 1000 / 1000
	if math.Abs(float64(ts-now)) > 10000 {
		t.Error("Invalid timestamp")
	}
}
