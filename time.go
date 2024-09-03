package kucoin

import (
	"context"
	"net/http"
)

type ServerTimeModel int64

// ServerTime returns the API server time.
func (as *ApiService) ServerTime(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	return as.Call(ctx, req)
}
