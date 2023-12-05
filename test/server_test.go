package test

import (
	"daemon_backend.bin/component/server"
	"fmt"
	"net/http"
	"testing"
)

// TestStartServer tests the successful start of the server and checks its health endpoint.
func TestStartServer(t *testing.T) {
	// Start the server in a goroutine
	go server.StartServer()

	// Wait for the server to be reachable with a maximum of 10 retries
	retries, err := server.WaitForServer("http://127.0.0.1:8080/", 10)
	if err != nil {
		t.Fatal(err)
	}

	// Log the successful server start information
	t.Log(fmt.Sprintf("Server started after %d tries", retries))

	// Define parameters for an HTTP GET request to the health endpoint
	method := "GET"
	endpoint := "http://127.0.0.1:8080/api/v2/healthz"
	payload := map[string]string{}
	expectedStatusCode := http.StatusOK
	expectedResponse := `{}`

	// Perform the HTTP request and validate the response
	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}
