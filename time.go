package kucoin

import (
	"net/http"
)

func (as ApiService) ServerTime() (int64, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	rsp, err := as.Call(req)
	if err != nil {
		return 0, err
	}
	var v int64
	rsp.ReadData(&v)
	return v, nil
}
