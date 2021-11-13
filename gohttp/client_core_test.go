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
	requestHeaders.Set("X-Request-Id", "TEST-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3")
	}

	if finalHeaders.Get("X-Request-Id") != "TEST-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}

}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no body expected when passing nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if string(body) != `["one", "two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

	})

}

func TestGetRequestBodyNilBody(t *testing.T) {
	client := httpClient{}

	body, err := client.getRequestBody("", nil)

	if err != nil {
		t.Error("no error expected when passing nil body")
	}

	if body != nil {
		t.Error("no body expected when passing nil body")
	}
}
