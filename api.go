package kucoin

import "log"

var (
	BaseURL = "https://openapi-v2.kucoin.com"
	Api     = &ApiService{
		requester: &BasicRequester{},
		signer:    &Sha256Signer{},
	}
)

type ApiService struct {
	requester Requester
	signer    Signer
}

func (k *ApiService) Call(request *Request) (*ApiResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[Recovery] panic recovered:", err)
		}
	}()
	request.Header.Set("Content-Type", "application/json")
	rsp, err := k.requester.Request(request, request.Timeout)
	if err != nil {
		return nil, err
	}
	return rsp.ApiResponse()
}
