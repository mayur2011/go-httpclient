package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalheaders := client.getRequestHeaders(requestHeaders)

	//Validation
	if len(finalheaders) != 3 {
		t.Error("expecting 3 headers")
	}

	if finalheaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content-type value")
	}

	if finalheaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user-agent value")
	}

	if finalheaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid x-request-id value")
	}
}
