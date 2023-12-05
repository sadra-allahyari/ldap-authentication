package extractor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ParseRequestBody parses the JSON request body of an HTTP request into a map[string]string.
// It takes an *http.Request as a parameter and returns a map[string]string representing the parsed data.
// If there is an error during parsing, it returns an error.
func ParseRequestBody(r *http.Request, requiredFields []string) (map[string]string, error) {
	var requestBody map[string]string

	// Decode the JSON request body into the map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, fmt.Errorf("error parsing request body: %s", err)
	}

	// Check for the presence of required fields
	var missingFields []string
	for _, field := range requiredFields {
		if _, ok := requestBody[field]; !ok {
			missingFields = append(missingFields, field)
		}
	}

	// If there are missing fields, return an error
	if len(missingFields) > 0 {
		missingFieldsStr := strings.Join(missingFields, ", ")
		return nil, fmt.Errorf("missing fields: %s", missingFieldsStr)
	}

	// Return the parsed data
	return requestBody, nil
}
