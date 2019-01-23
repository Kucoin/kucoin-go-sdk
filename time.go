package kucoin

import (
	"net/http"
)

func (as *ApiService) ServerTime() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	return as.call(req)
}
