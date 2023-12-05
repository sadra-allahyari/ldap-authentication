package handler

import (
	"context"
	"daemon_backend.bin/component/auth"
	"daemon_backend.bin/component/database"
	"daemon_backend.bin/component/extractor"
	"daemon_backend.bin/component/ldap"
	"daemon_backend.bin/component/respond"
	"daemon_backend.bin/source/update/database/data_store/user/userDB"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// LoginHandler handles the login request.
// It expects a JSON request body containing 'username' and 'password' fields.
// The function performs the following steps:
// 1. Parse the JSON request body.
// 2. Obtain user credentials from the parsed JSON.
// 3. Establish an LDAP connection using the bind function.
// 4. Authenticate the user using the LDAP connection.
// 5. Respond with an error if authentication fails, including handling LDAP-specific errors.
// 6. If authentication is successful, respond with the API key.
// 7: Check if the user already exists.
// 8: User does not exist, create a new user.
// 9: Retrieve the API key for the created user.
// 10: Respond with the API key
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Step 0: Specify the required fields for JSON request body
	requiredFields := []string{"username", "password"}

	// Step 1: Parse JSON body
	requestBody, err := extractor.ParseRequestBody(r, requiredFields)
	if err != nil {
		// Convert newline-separated error message to JSON format and respond with a bad request
		errorMessage := strings.ReplaceAll(err.Error(), "\n", "\",\n\"")
		respond.ErrRespond(w, http.StatusBadRequest, errorMessage)
		return
	}

	// Step 2: Get user credentials from JSON body
	username := requestBody["username"]
	password := requestBody["password"]

	// Step 3: Use the bind function to obtain the LDAP connection
	ldapConn, err := ldap.Bind()
	if err != nil {
		// Respond with an internal server error if LDAP connection fails
		respond.ErrRespond(w, http.StatusInternalServerError, fmt.Sprintf("LDAP connection error: %s", err))
		return
	}

	// Step 4: Authenticate the user
	err = auth.AuthenticateUser(ldapConn, username, password)
	if err != nil {
		// Step 5: Handle LDAP-specific errors
		ldapErrCode, addsErrCode, extractErr := extractor.ExtractErrorCode(err.Error())
		if extractErr != nil {
			// Respond with an unauthorized status and an error message if extracting LDAP error code fails
			respond.ErrRespond(w, http.StatusUnauthorized, fmt.Sprintf("Error extracting LDAP error code: %s", extractErr))
			return
		}

		// Use handleLdapError function to respond with the appropriate LDAP error message
		handleLdapError(w, ldapErrCode, addsErrCode)
		return
	}

	// Step 6: If authentication is successful, respond with the API key
	dn := auth.GetUserDN(username, extractor.ExtractStrFromFile("ldap", "baseDN"))

	conn, err := database.DbConnector()
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error connecting to the database"})
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
			respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error closing the database connection"})
		}
	}()

	query := userDB.New(conn)

	// Declare apiKey outside the if block
	var apiKey string

	// Step 7: Check if the user already exists
	apiKey, err = query.GetAPIKey(context.TODO(), dn)
	if err == nil {
		// User already exists, return early without creating a new user
		respond.JsonRespond(w, http.StatusOK, map[string]interface{}{"key": apiKey})
		return
	}

	// Step 8: User does not exist, create a new user
	_, err = query.CreateUser(context.TODO(), dn)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error creating user"})
		return
	}

	// Step 9: Retrieve the API key for the created user
	apiKey, err = query.GetAPIKey(context.TODO(), dn)
	if err != nil {
		log.Printf("Error getting API key: %v", err)
		respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error getting API key"})
		return
	}

	// Step 10: Respond with the API key
	respond.JsonRespond(w, http.StatusOK, map[string]interface{}{"key": apiKey})
}
