package kucoin

import (
	"log"
	"os"
)

type ApiService struct {
	ApiBaseURI         string
	InsecureSkipVerify bool
	requester          Requester
	signer             Signer
}

const ApiBaseURI = "https://openapi-v2.kucoin.com"

func NewPublicApi() *ApiService {
	as := &ApiService{
		ApiBaseURI: ApiBaseURI,
		requester:  &BasicRequester{},
	}
	return as
}

func NewPrivateApi(key, secret, passphrase string) *ApiService {
	as := NewPublicApi()
	as.signer = NewKcSigner(key, secret, passphrase)
	return as
}

func NewPublicApiFromEnv() *ApiService {
	as := NewPublicApi()
	if u := os.Getenv("API_BASE_URI"); u != "" {
		as.ApiBaseURI = u
	}
	return as
}

func NewPrivateApiFromEnv() *ApiService {
	as := NewPrivateApi(os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("API_PASSPHRASE"))
	if u := os.Getenv("API_BASE_URI"); u != "" {
		as.ApiBaseURI = u
	}
	return as
}

func (as *ApiService) Call(request *Request) (*ApiResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[Recovery] panic recovered:", err)
		}
	}()
	request.BaseURL = as.ApiBaseURI
	request.InsecureSkipVerify = as.InsecureSkipVerify
	request.Header.Set("Content-Type", "application/json")
	if as.signer != nil {
		// todo
		request.Header.Set("KC-API-KEY", "")
		request.Header.Set("KC-API-SIGN", "")
		request.Header.Set("KC-API-TIMESTAMP", "")
		request.Header.Set("KC-API-PASSPHRASE", "")
	}
	rsp, err := as.requester.Request(request, request.Timeout)
	if err != nil {
		return nil, err
	}
	return rsp.ApiResponse()
}
