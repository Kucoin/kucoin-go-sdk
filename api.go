package kucoin

import (
	"net/http"
	"time"
)

var (
	BaseURL = "https://openapi-v2.kucoin.com"
	Api     = &KcApiService{
		requester: &BasicRequester{},
		signer:    &Sha256Signer{},
	}
)

type ApiService interface {
	Call(request *Request, timeout time.Duration) (*http.Response, error)
}

type KcApiService struct {
	requester Requester
	signer    Signer
}

func (k *KcApiService) call(request *Request) (*Response, error) {
	rsp, err := k.requester.Request(request, request.Timeout)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (k *KcApiService) CallApi(request *Request) (*ApiResponse, error) {
	request.Header.Set("Content-Type", "application/json")
	rsp, err := k.call(request)
	if err != nil {
		return nil, err
	}
	return rsp.ApiResponse()
}
