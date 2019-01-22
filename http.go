package kucoin

import (
	"bytes"
	"crypto/tls"
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
	fullURL            string
	requestURI         string
	BaseURI            string
	Method             string
	Path               string
	Query              url.Values
	Body               io.Reader
	Header             http.Header
	Timeout            time.Duration
	InsecureSkipVerify bool
}

func NewRequest(method, path string, params map[string]string) *Request {
	r := &Request{
		Method: method,
		Path:   path,
	}
	if r.Path == "" {
		r.Path = "/"
	}
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	r.Query = make(url.Values)
	r.Header = make(http.Header)
	r.Body = bytes.NewBufferString("")
	r.addParams(params)
	r.Timeout = 30 * time.Second
	return r
}

func (r *Request) addParams(params map[string]string) {
	switch r.Method {
	case http.MethodGet, http.MethodDelete:
		for key, value := range params {
			r.Query.Add(key, value)
		}
	default:
		q := &url.Values{}
		for key, value := range params {
			q.Add(key, value)
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
		r.fullURL = fmt.Sprintf("%s%s", r.BaseURI, r.Path)
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

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

type Requester interface {
	Request(request *Request, timeout time.Duration) (*Response, error)
}

type BasicRequester struct {
}

func (bh *BasicRequester) Request(request *Request, timeout time.Duration) (*Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: request.InsecureSkipVerify},
	}
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
		body:     nil,
	}, nil
}

type Response struct {
	request *Request
	*http.Response
	body []byte
}

func (r *Response) ReadBody() ([]byte, error) {
	if r.body == nil {
		r.body = make([]byte, 0)
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
