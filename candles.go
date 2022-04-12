package kucoin

import (
	"net/http"
)

// CandleHistory returns the klines of the specified symbol between two points in time
func (as *ApiService) CandleHistory(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/market/candles", params)
	return as.Call(req)
}

type CandleProperty string

type Candle []*CandleProperty
