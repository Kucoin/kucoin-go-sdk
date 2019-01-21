package kucoin

import (
	"net/http"
)

type Time struct {
}

type Timestamp struct {
	Data int64 `json:"data"`
}

func (t *Time) Timestamp() (int64, error) {
	req := NewRequest(http.MethodGet, "/api/v1/timestamp", map[string]string{})
	rsp, err := Api.CallApi(req)
	if err != nil {
		return 0, err
	}
	v := &Timestamp{}
	rsp.ApiData(v)
	return v.Data, nil
}
