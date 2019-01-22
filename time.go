package kucoin

import (
	"net/http"
)

func (as ApiService) ServerTime(v *int64) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return rsp, err
	}
	if err := rsp.ReadData(v); err != nil {
		return rsp, err
	}
	return rsp, nil
}
