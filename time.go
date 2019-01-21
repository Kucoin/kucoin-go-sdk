package kucoin

import (
	"net/http"
)

func Timestamp() (int64, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", nil)
	rsp, err := Api.Call(req)
	if err != nil {
		return 0, err
	}
	type Data struct {
		Data int64 `json:"data"`
	}
	v := &Data{}
	rsp.ApiData(v)
	return v.Data, nil
}
