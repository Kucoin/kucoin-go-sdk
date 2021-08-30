package kucoin

import (
	"net/http"
)

type ServerTimeModel int64

// ServerTime returns the API server time.
func (as *ApiService) ServerTime() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	return as.Call(req)
}
