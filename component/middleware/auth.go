package middleware

import (
	"context"
	"daemon_backend.bin/component/database"
	"daemon_backend.bin/component/respond"
	"daemon_backend.bin/source/update/database/data_store/user/userDB"
	"log"
	"net/http"
)

// AuthMiddleware is a middleware that checks for the presence and validity of an API key in the HTTP header.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Step 1: Get the API key from the Authorization header
		apiKey := r.Header.Get("Authorization")

		// Step 2: Check if the API key is present and not empty
		if apiKey != "" {
			// Step 3: Validate the API key against the database
			if isValidAPIKey(apiKey) {
				// Step 4: Call the next handler in the chain
				next.ServeHTTP(w, r)
				return
			}
		}

		// Step 5: If the API key is missing or empty, respond with an error
		respond.ErrRespond(w, http.StatusUnauthorized, "API key is missing or invalid")
	})
}

// isValidAPIKey checks if the provided API key is valid by looking it up in the database.
// It returns true if the API key is found in the database, indicating its validity.
// If there is an error during the database lookup or connection, it returns false.
func isValidAPIKey(apiKey string) bool {
	// Connect to the database
	conn, err := database.DbConnector()
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		return false
	}
	defer func() {
		// Close the database connection when the function exits
		if err := conn.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		}
	}()

	// Create a userDB query instance
	query := userDB.New(conn)

	// Check if the API key exists in the database
	userID, err := query.GetUserIDByAPIKey(context.TODO(), apiKey)
	if err != nil {
		log.Printf("Error checking API key in the database: %v", err)
		return false
	}

	// If a user is found with the given API key, consider it valid
	return userID != 0
}
