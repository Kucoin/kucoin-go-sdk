package kucoin

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// A Request represents a HTTP request.
type Request struct {
	fullURL       string
	requestURI    string
	BaseURI       string
	Method        string
	Path          string
	Query         url.Values
	Body          []byte
	Header        http.Header
	Timeout       time.Duration
	SkipVerifyTls bool
}

// NewRequest creates a instance of Request.
func NewRequest(method, path string, params interface{}) *Request {
	r := &Request{
		Method:  method,
		Path:    path,
		Query:   make(url.Values),
		Header:  make(http.Header),
		Body:    []byte{},
		Timeout: 30 * time.Second,
	}
	if r.Path == "" {
		r.Path = "/"
	}
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	r.addParams(params)
	return r
}

func (r *Request) addParams(p interface{}) {
	if p == nil {
		return
	}
	switch r.Method {
	case http.MethodGet, http.MethodDelete:
		for key, value := range p.(map[string]string) {
			r.Query.Add(key, value)
		}
	default:
		b, err := json.Marshal(p)
		if err != nil {
			log.Panic("Cannot marshal params to JSON string:", err.Error())
		}
		r.Body = b
	}
}

// RequestURI returns the request uri.
func (r *Request) RequestURI() string {
	if r.requestURI != "" {
		return r.requestURI
	}

	fu := r.FullURL()
	u, err := url.Parse(fu)
	if err != nil {
		r.requestURI = "/"
	} else {
		r.requestURI = u.RequestURI()
	}
	return r.requestURI
}

// FullURL returns the full url.
func (r *Request) FullURL() string {
	if r.fullURL != "" {
		return r.fullURL
	}
	r.fullURL = fmt.Sprintf("%s%s", r.BaseURI, r.Path)
	if len(r.Query) > 0 {
		if strings.Contains(r.fullURL, "?") {
			r.fullURL += "&" + r.Query.Encode()
		} else {
			r.fullURL += "?" + r.Query.Encode()
		}
	}
	return r.fullURL
}

// HttpRequest creates a instance of *http.Request.
func (r *Request) HttpRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.FullURL(), bytes.NewBuffer(r.Body))
	if err != nil {
		return nil, err
	}

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

// Requester contains Request() method, can launch a http request.
type Requester interface {
	Request(request *Request, timeout time.Duration) (*Response, error)
}

// A BasicRequester represents a basic implement of Requester by http.Client.
type BasicRequester struct {
}

// Request makes a http request.
func (br *BasicRequester) Request(request *Request, timeout time.Duration) (*Response, error) {
	tr := http.DefaultTransport
	tc := tr.(*http.Transport).TLSClientConfig
	if tc == nil {
		tc = &tls.Config{InsecureSkipVerify: request.SkipVerifyTls}
	} else {
		tc.InsecureSkipVerify = request.SkipVerifyTls
	}

	cli := http.DefaultClient
	cli.Transport, cli.Timeout = tr, timeout

	req, err := request.HttpRequest()
	if err != nil {
		return nil, err
	}
	// Prevent re-use of TCP connections
	// req.Close = true

	rid := time.Now().UnixNano()

	if DebugMode {
		dump, _ := httputil.DumpRequest(req, true)
		logrus.Debugf("Sent a HTTP request#%d: %s", rid, string(dump))
	}

	rsp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if DebugMode {
		dump, _ := httputil.DumpResponse(rsp, true)
		logrus.Debugf("Received a HTTP response#%d: %s", rid, string(dump))
	}

	return NewResponse(
		request,
		rsp,
		nil,
	), nil
}

// A Response represents a HTTP response.
type Response struct {
	request *Request
	*http.Response
	body []byte
}

// NewResponse Creates a new Response
func NewResponse(
	request *Request,
	response *http.Response,
	body []byte,
) *Response {
	return &Response{
		request:  request,
		Response: response,
		body:     body,
	}
}

// ReadBody read the response data, then return it.
func (r *Response) ReadBody() ([]byte, error) {
	if r.body != nil {
		return r.body, nil
	}

	r.body = make([]byte, 0)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.body = b
	return r.body, nil
}

// ReadJsonBody read the response data as JSON into v.
func (r *Response) ReadJsonBody(v interface{}) error {
	b, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// The predefined API codes
const (
	ApiSuccess = "200000"
)

// An ApiResponse represents a API response wrapped Response.
type ApiResponse struct {
	response *Response
	Code     string          `json:"code"`
	RawData  json.RawMessage `json:"data"` // delay parsing
	Message  string          `json:"msg"`
}

// HttpSuccessful judges the success of http.
func (ar *ApiResponse) HttpSuccessful() bool {
	return ar.response.StatusCode == http.StatusOK
}

// ApiSuccessful judges the success of API.
func (ar *ApiResponse) ApiSuccessful() bool {
	return ar.Code == ApiSuccess
}

// ReadData read the api response `data` as JSON into v.
func (ar *ApiResponse) ReadData(v interface{}) error {
	if !ar.HttpSuccessful() {
		rsb, _ := ar.response.ReadBody()
		m := fmt.Sprintf("[HTTP]Failure: status code is NOT 200, %s %s with body=%s, respond code=%d body=%s",
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			string(ar.response.request.Body),
			ar.response.StatusCode,
			string(rsb),
		)
		return errors.New(m)
	}

	if !ar.ApiSuccessful() {
		m := fmt.Sprintf("[API]Failure: api code is NOT %s, %s %s with body=%s, respond code=%s message=\"%s\" data=%s",
			ApiSuccess,
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			string(ar.response.request.Body),
			ar.Code,
			ar.Message,
			string(ar.RawData),
		)
		return errors.New(m)
	}
	// when input parameter v is nil, read nothing and return nil
	if v == nil {
		return nil
	}

	if len(ar.RawData) == 0 {
		m := fmt.Sprintf("[API]Failure: try to read empty data, %s %s with body=%s, respond code=%s message=\"%s\" data=%s",
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			string(ar.response.request.Body),
			ar.Code,
			ar.Message,
			string(ar.RawData),
		)
		return errors.New(m)
	}

	return json.Unmarshal(ar.RawData, v)
}

// ReadPaginationData read the data `items` as JSON into v, and returns *PaginationModel.
func (ar *ApiResponse) ReadPaginationData(v interface{}) (*PaginationModel, error) {
	p := &PaginationModel{}
	if err := ar.ReadData(p); err != nil {
		return nil, err
	}
	if err := p.ReadItems(v); err != nil {
		return p, err
	}
	return p, nil
}
