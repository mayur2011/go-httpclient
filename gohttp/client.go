package gohttp

import "net/http"

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
}

//common header setup which will always be there besides headers from client request
type httpClient struct {
	Headers http.Header
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}
