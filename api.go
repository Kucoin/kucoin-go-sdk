package kucoin

import (
	"bytes"
	"log"
	"os"
)

type ApiService struct {
	apiBaseURI    string
	apiKey        string
	apiSecret     string
	apiPassphrase string
	requester     Requester
	signer        Signer
	SkipVerifyTls bool
}

const ApiBaseURI = "https://openapi-v2.kucoin.com"

type ApiServiceOption func(service *ApiService)

func ApiBaseURIOption(uri string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiBaseURI = uri
	}
}

func ApiKeyOption(key string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiKey = key
	}
}

func ApiSecretOption(secret string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiSecret = secret
	}
}

func ApiPassPhraseOption(passPhrase string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiPassphrase = passPhrase
	}
}

func NewApiService(opts ...ApiServiceOption) *ApiService {
	as := &ApiService{
		requester: &BasicRequester{},
	}
	for _, opt := range opts {
		opt(as)
	}
	if as.apiBaseURI == "" {
		as.apiBaseURI = ApiBaseURI
	}
	if as.apiKey != "" {
		as.signer = NewKcSigner(as.apiKey, as.apiSecret, as.apiPassphrase)
	}
	return as
}

func NewApiServiceFromEnv() *ApiService {
	return NewApiService(
		ApiBaseURIOption(os.Getenv("API_BASE_URI")),
		ApiKeyOption(os.Getenv("API_KEY")),
		ApiSecretOption(os.Getenv("API_SECRET")),
		ApiPassPhraseOption(os.Getenv("API_PASSPHRASE")),
	)
}

func (as *ApiService) call(request *Request) (*ApiResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[Recovery] panic recovered:", err)
		}
	}()

	request.BaseURI = as.apiBaseURI
	request.SkipVerifyTls = as.SkipVerifyTls
	request.Header.Set("Content-Type", "application/json")
	if as.signer != nil {
		var b bytes.Buffer
		b.WriteString(request.Method)
		b.WriteString(request.RequestURI())
		b.Write(request.Body)
		h := as.signer.(*KcSigner).Headers(b.String())
		for k, v := range h {
			request.Header.Set(k, v)
		}
	}

	rsp, err := as.requester.Request(request, request.Timeout)
	if err != nil {
		return nil, err
	}

	ar := &ApiResponse{response: rsp}
	if err := rsp.ReadJsonBody(ar); err != nil {
		return nil, err
	}
	return ar, nil
}
