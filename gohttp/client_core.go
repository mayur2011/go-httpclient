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

	for header, value := range headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}

	return client.Do(request)
}
