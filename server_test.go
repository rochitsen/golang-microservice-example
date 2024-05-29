package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	// arrange
	server := httptest.NewServer(http.HandlerFunc(handler))
	
	// act
	actualResponse, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	// assert HTTP status code
	if actualResponse.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", actualResponse.StatusCode)
	}

	respBody, err := io.ReadAll(actualResponse.Body)
	if err != nil {
		t.Fatal(err)
	}
	
	var actual map[string]interface{}
	if err := json.Unmarshal(respBody, &actual); err != nil {
		t.Fatal(err)
	}

	expectedResponse := `{
		"name": "demo",
		"author": "rochitsen"
	}`
	
	var expected map[string]interface{}
	if err := json.Unmarshal([]byte(expectedResponse), &expected); err != nil {
		t.Fatal(err)
	}

	// assert response
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	// cleanup
	server.Close()
}