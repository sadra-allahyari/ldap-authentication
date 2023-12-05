package server

import (
	"fmt"
	"net/http"
	"time"
)

// WaitForServer waits for the server at the specified URL to become reachable,
// making HTTP GET requests with a maximum number of retries.
// It returns the number of retries made and an error if the server is not reachable
// after the specified number of retries.
func WaitForServer(url string, maxRetries int) (int, error) {
	// Initialize the retry counter
	tries := 0

	// Loop until the maximum number of retries is reached
	for i := 0; i < maxRetries; i++ {
		// Attempt to make an HTTP GET request to the server
		resp, err := http.Get(url)

		// Increment the retry counter
		tries++
		fmt.Println(tries)

		// Check if the request was successful and the server returned a "Not Found" status
		if err == nil && resp.StatusCode == http.StatusNotFound {
			// Server is reachable, close the response body and break out of the loop
			err := resp.Body.Close()
			if err != nil {
				return tries, err
			}
			return tries, nil
		}

		// Sleep for a short duration before retrying
		time.Sleep(1 * time.Millisecond)
	}

	// Return an error if the server is not reachable after the specified number of retries
	return tries, fmt.Errorf("unable to reach server after %d tries", maxRetries)
}
