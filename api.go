/*
Package kucoin provides two kinds of APIs: `RESTful API` and `WebSocket feed`.
The official document: https://docs.kucoin.com
*/
package kucoin

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// Debug mode will record API and WebSocket logs to files in the directory "kucoin.LogDirectory" according to the minimum log level "kucoin.LogLevel".
	DebugMode = os.Getenv("API_DEBUG_MODE") == "1"
	// LogLevel is the lowest logging level of logrus, the default value is logrus.DebugLevel.
	LogLevel = logrus.DebugLevel
	// LogDirectory is the directory of log file, the default value is "/tmp".
	LogDirectory = "/tmp"
)

func init() {
	// Initialize the logging component
	logFile := fmt.Sprintf("%s/kucoin-sdk-%s.log", LogDirectory, time.Now().Format("2006-01-02"))
	logWriter, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicf("Open file failed: %s", err.Error())
	}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(logWriter)
	logrus.SetLevel(LogLevel)
}

// An ApiService provides a HTTP client and a signer to make a HTTP request with the signature to KuCoin API.
type ApiService struct {
	apiBaseURI       string
	apiKey           string
	apiSecret        string
	apiPassphrase    string
	apiSkipVerifyTls bool
	apiDebugMode     bool
	requester        Requester
	signer           Signer
}

// ProductionApiBaseURI is api base uri for production.
const ProductionApiBaseURI = "https://openapi-v2.kucoin.com"

// An ApiServiceOption is a option parameter to create the instance of ApiService.
type ApiServiceOption func(service *ApiService)

// ApiBaseURIOption creates a instance of ApiServiceOption about apiBaseURI.
func ApiBaseURIOption(uri string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiBaseURI = uri
	}
}

// ApiKeyOption creates a instance of ApiServiceOption about apiKey.
func ApiKeyOption(key string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiKey = key
	}
}

// ApiSecretOption creates a instance of ApiServiceOption about apiSecret.
func ApiSecretOption(secret string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiSecret = secret
	}
}

// ApiPassPhraseOption creates a instance of ApiServiceOption about apiPassPhrase.
func ApiPassPhraseOption(passPhrase string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiPassphrase = passPhrase
	}
}

// ApiSkipVerifyTlsOption creates a instance of ApiServiceOption about apiSkipVerifyTls.
func ApiSkipVerifyTlsOption(skipVerifyTls bool) ApiServiceOption {
	return func(service *ApiService) {
		service.apiSkipVerifyTls = skipVerifyTls
	}
}

// NewApiService creates a instance of ApiService by passing ApiServiceOptions, then you can call methods.
func NewApiService(opts ...ApiServiceOption) *ApiService {
	as := &ApiService{requester: &BasicRequester{}}
	for _, opt := range opts {
		opt(as)
	}
	if as.apiBaseURI == "" {
		as.apiBaseURI = ProductionApiBaseURI
	}
	if as.apiKey != "" {
		as.signer = NewKcSigner(as.apiKey, as.apiSecret, as.apiPassphrase)
	}
	return as
}

// NewApiServiceFromEnv creates a instance of ApiService by environmental variables such as `API_BASE_URI` `API_KEY` `API_SECRET` `API_PASSPHRASE`, then you can call the methods of ApiService.
func NewApiServiceFromEnv() *ApiService {
	return NewApiService(
		ApiBaseURIOption(os.Getenv("API_BASE_URI")),
		ApiKeyOption(os.Getenv("API_KEY")),
		ApiSecretOption(os.Getenv("API_SECRET")),
		ApiPassPhraseOption(os.Getenv("API_PASSPHRASE")),
		ApiSkipVerifyTlsOption(os.Getenv("API_SKIP_VERIFY_TLS") == "1"),
	)
}

// Call calls the API by passing *Request and returns *ApiResponse.
func (as *ApiService) Call(request *Request) (*ApiResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[Recovery] panic recovered:", err)
		}
	}()

	request.BaseURI = as.apiBaseURI
	request.SkipVerifyTls = as.apiSkipVerifyTls
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
		rb, _ := rsp.ReadBody()
		m := fmt.Sprintf("[Parse]Failure: parse JSON body failed because %s, %s %s with body=%s, respond code=%d body=%s",
			err.Error(),
			rsp.request.Method,
			rsp.request.RequestURI(),
			string(rsp.request.Body),
			rsp.StatusCode,
			string(rb),
		)
		return ar, errors.New(m)
	}
	return ar, nil
}
