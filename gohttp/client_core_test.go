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

func TestGetRequestBodyAsNilBody(t *testing.T) {
	client := httpClient{}
	body, err := client.getRequestBody("application/json", nil)
	if err != nil {
		t.Error("no error expected with nil body request")
	}

	if body != nil {
		t.Error("no error expected with nil body request")
	}
}
func TestGetRequestBodyAsJSON(t *testing.T) {
	client := httpClient{}
	requestBody := []string{"first", "last"}
	body, err := client.getRequestBody("application/json", requestBody)

	if err != nil {
		t.Error("not expected error with marshaling slice as json")
	}

	if string(body) != `["first", "last"]` {
		t.Error("not expected error invalid request body")
	}

}
func TestGetRequestBodyAsXML(t *testing.T) {

}
func TestGetRequestBodyJSONAsDefault(t *testing.T) {

}
