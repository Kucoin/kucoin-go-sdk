package kucoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	timeout    time.Duration
	fullURL    string
	requestURI string
	BaseURL    string
	Method     string
	Path       string
	Query      url.Values
	Body       io.Reader
	Headers    map[string]string
}

func NewRequest(method, path string, params map[string]interface{}) *Request {
	r := &Request{
		Method: method,
		Path:   path,
	}
	r.BaseURL = BaseURL
	if r.Path == "" {
		r.Path = "/"
	}
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	r.addParams(params)
	return r
}

func (r *Request) addParams(params map[string]interface{}) {
	switch r.Method {
	case http.MethodGet, http.MethodDelete:
		for key, value := range params {
			r.Query.Add(key, value.(string))
		}
	default:
		q := &url.Values{}
		for key, value := range params {
			q.Add(key, value.(string))
		}
		b, err := json.Marshal(params)
		if err != nil {
			log.Panic("Cannot marshal params to JSON string:", err.Error())
		}
		r.Body = bytes.NewBuffer(b)
	}
}

func (r *Request) RequestURI() string {
	if r.requestURI == "" {
		fu := r.FullURL()
		u, err := url.Parse(fu)
		if err != nil {
			r.requestURI = "/"
		} else {
			r.requestURI = u.RequestURI()
		}
	}
	return r.requestURI
}

func (r *Request) FullURL() string {
	if r.fullURL == "" {
		r.fullURL = fmt.Sprintf("%s%s", r.BaseURL, r.Path)
		if len(r.Query) > 0 {
			if strings.Contains(r.fullURL, "?") {
				r.fullURL += "&" + r.Query.Encode()
			} else {
				r.fullURL += "?" + r.Query.Encode()
			}
		}
	}
	return r.fullURL
}

func (r *Request) HttpRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.FullURL(), r.Body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// Add Queries
	q := req.URL.Query()
	for key, values := range r.Query {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

type Requester interface {
	Request(request *Request, timeout time.Duration) (*Response, error)
}

type BasicRequester struct {
}

func (bh *BasicRequester) Request(request *Request, timeout time.Duration) (*Response, error) {
	tr := &http.Transport{}
	cli := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	req, err := request.HttpRequest()
	if err != nil {
		return nil, err
	}

	rsp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{
		request:  request,
		Response: rsp,
	}, nil
}

type Response struct {
	request *Request
	*http.Response
	body []byte
}

func (r *Response) ReadBody() ([]byte, error) {
	if len(r.body) == 0 {
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.body = b
	}
	return r.body, nil
}

func (r *Response) ReadJsonBody(v interface{}) error {
	b, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (r *Response) ApiResponse() (*ApiResponse, error) {
	ar := &ApiResponse{response: r}
	if err := r.ReadJsonBody(ar); err != nil {
		return nil, err
	}
	return ar, nil
}

const (
	ApiSuccess = "200000"
)

type ApiResponse struct {
	response *Response
	Code     string `json:"code"`
	Message  string `json:"msg"`
}

func (ar *ApiResponse) MustBeSuccessful() {
	if ar.response.StatusCode != http.StatusOK {
		rb, _ := ar.response.ReadBody()
		log.Panicf("[HTTP]Failure: status code is NOT 200, %s %s with body=%s, respond code=%d body=%s",
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			"todo",
			ar.response.StatusCode,
			string(rb),
		)
	}

	if ar.Code != ApiSuccess {
		rb, _ := ar.response.ReadBody()
		log.Panicf("[API]Failure: api code is NOT %s, %s %s with body=%s, respond code=%s message=\"%s\" data=%s",
			ApiSuccess,
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			"todo",
			ar.Code,
			ar.Message,
			string(rb),
		)
	}
}

func (ar *ApiResponse) ApiData(v interface{}) {
	ar.MustBeSuccessful()
	err := ar.response.ReadJsonBody(v)
	if err != nil {
		log.Panicf("[API]Failure: Parse data failed: %s, %s %s with body=%s, respond code=%d",
			err.Error(),
			ar.response.request.Method,
			ar.response.request.RequestURI(),
			"todo",
			ar.response.StatusCode,
		)
	}
}
