package gohttp

import (
	"errors"
	"net/http"
)

func (C *httpClient) do(method string, url string, headers http.Header, req interface{}) (*http.Response, error) {

	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}
	return client.Do(request)
}