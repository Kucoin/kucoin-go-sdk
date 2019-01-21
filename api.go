package kucoin

import "log"

var (
	BaseURL       = "https://openapi-v2.kucoin.com"
	ApiKey        = "your api key"
	ApiSecret     = "your api secret"
	ApiPassphrase = "your api passphrase"
	PublicApi     = &ApiService{
		requester: &BasicRequester{},
	}
	PrivateApi = &ApiService{
		requester: &BasicRequester{},
		signer:    NewKcSigner(ApiKey, ApiSecret, ApiPassphrase),
	}
)

type ApiService struct {
	requester Requester
	signer    Signer
}

func (as *ApiService) Call(request *Request) (*ApiResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[Recovery] panic recovered:", err)
		}
	}()
	request.BaseURL = BaseURL
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
