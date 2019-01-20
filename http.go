package kucoin

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	fullURL    string
	requestURI string
	BaseURL    string
	Timeout    time.Duration
	Method     string
	URL        string
	Query      url.Values
	Body       io.Reader
	Headers    map[string]string
}

func (r *Request) FullURL() string {
	if r.fullURL == "" {
		r.fullURL = fmt.Sprintf("%s%s", r.BaseURL, r.URL)
	}
	return r.fullURL
}

func (r *Request) RequestURI() string {
	if r.requestURI == "" {
		fu := r.FullURL()
		u, err := url.Parse(fu)
		if err != nil {
			return "/"
		}
		r.requestURI = u.RequestURI()
	}
	return r.requestURI
}

type Http struct {
}

func (h *Http) Request(request *Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.URL, request.Body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	// Add Queries
	q := req.URL.Query()
	for key, values := range request.Query {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	req.URL.RawQuery = q.Encode()

	tr := &http.Transport{}
	cli := &http.Client{
		Transport: tr,
		Timeout:   request.Timeout,
	}
	rsp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	return rsp, nil

}
