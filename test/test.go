package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// performHTTPRequest is a helper function for testing HTTP requests.
// It takes in the testing object, HTTP method, endpoint, payload, expected status code,
// and expected response. It performs the HTTP request and validates the response.
func performHTTPRequest(t *testing.T, method, endpoint string, payload map[string]string, expectedStatusCode int, expectedResponse string) {
	var req *http.Request
	var err error

	// Create an HTTP request based on the specified method
	switch method {
	case "GET":
		req, err = http.NewRequest("GET", endpoint, nil)
	case "POST":
		// Convert the payload to a JSON string
		jsonData, err := json.Marshal(payload)
		if err != nil {
			t.Fatal("Error encoding JSON:", err)
		}
		req, err = http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
	default:
		t.Fatalf("Unsupported HTTP method: %s", method)
	}

	if err != nil {
		t.Fatal(err)
	}

	// Perform the HTTP request to the server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Error making HTTP request: %v", err)
	}
	defer func() {
		// Close the response body, and handle any errors
		err := resp.Body.Close()
		if err != nil {
			t.Fatalf("Error closing response body: %v", err)
		}
	}()

	// Check the HTTP status code
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Compare the actual response with the expected response
	expected := expectedResponse
	if got := string(body); got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
