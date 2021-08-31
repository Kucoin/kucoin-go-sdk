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
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// Version is SDK version.
	Version = "1.2.10"
	// DebugMode will record the logs of API and WebSocket to files in the directory "kucoin.LogDirectory" according to the minimum log level "kucoin.LogLevel".
	DebugMode = os.Getenv("API_DEBUG_MODE") == "1"
)

func init() {
	// Initialize the logging component by default
	logrus.SetLevel(logrus.DebugLevel)
	if runtime.GOOS == "windows" {
		SetLoggerDirectory("tmp")
	} else {
		SetLoggerDirectory("/tmp")
	}
}

// SetLoggerDirectory sets the directory for logrus output.
func SetLoggerDirectory(directory string) {
	var logFile string
	if !DebugMode {
		logFile = os.DevNull
	} else {
		logFile = fmt.Sprintf("%s/kucoin-sdk-%s.log", directory, time.Now().Format("2006-01-02"))
	}
	logWriter, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Panicf("Open file failed: %s", err.Error())
	}
	logrus.SetOutput(logWriter)
}

// An ApiService provides a HTTP client and a signer to make a HTTP request with the signature to KuCoin API.
type ApiService struct {
	apiBaseURI       string
	apiKey           string
	apiSecret        string
	apiPassphrase    string
	apiSkipVerifyTls bool
	requester        Requester
	signer           Signer
	apiKeyVersion    string
}

// ProductionApiBaseURI is api base uri for production.
const ProductionApiBaseURI = "https://api.kucoin.com"

/**
	Note: about api key version
    To reinforce the security of the API, KuCoin upgraded the API key to version 2.0, the validation logic has also been changed. It is recommended to create(https://www.kucoin.com/account/api) and update your API key to version 2.0. The API key of version 1.0 will be still valid until May 1, 2021.
*/

// ApiKeyVersionV1 is v1 api key version
const ApiKeyVersionV1 = "1"

// ApiKeyVersionV2 is v2 api key version
const ApiKeyVersionV2 = "2"

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

// ApiRequesterOption creates a instance of ApiServiceOption about requester.
func ApiRequesterOption(requester Requester) ApiServiceOption {
	return func(service *ApiService) {
		service.requester = requester
	}
}

// ApiKeyVersionOption creates a instance of ApiServiceOption about apiKeyVersion.
func ApiKeyVersionOption(apiKeyVersion string) ApiServiceOption {
	return func(service *ApiService) {
		service.apiKeyVersion = apiKeyVersion
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

	if as.apiKeyVersion == "" {
		as.apiKeyVersion = ApiKeyVersionV1
	}

	if as.apiKey != "" {
		if as.apiKeyVersion == ApiKeyVersionV1 {
			as.signer = NewKcSigner(as.apiKey, as.apiSecret, as.apiPassphrase)
		} else {
			as.signer = NewKcSignerV2(as.apiKey, as.apiSecret, as.apiPassphrase)
		}
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
		ApiKeyVersionOption(os.Getenv("API_KEY_VERSION")),
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
	request.Header.Set("User-Agent", "KuCoin-Go-SDK/"+Version)
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
