package handler

import (
	"context"
	"daemon_backend.bin/component/database"
	"daemon_backend.bin/component/respond"
	"daemon_backend.bin/source/update/database/data_store/user/userDB"
	"log"
	"net/http"
)

func TerminateHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Connect to the database
	conn, err := database.DbConnector()
	if err != nil {
		log.Printf("Error connecting to the database: %v", err)
		respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error connecting to the database"})
		return
	}
	defer func() {
		// Step 2: Close the database connection when done
		if err := conn.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
			respond.JsonRespond(w, http.StatusInternalServerError, map[string]interface{}{"error": "Error closing the database connection"})
		}
	}()

	// Step 3: Create a userDB query instance
	query := userDB.New(conn)

	// Step 4: Retrieve the API key from the Authorization header
	apiKey := r.Header.Get("Authorization")

	// Step 5: Delete the API key (terminate sessions) from the database
	err = query.DeleteAPIKey(context.TODO(), apiKey)
	if err != nil {
		// Step 6: Respond with an error if unable to delete other sessions
		respond.ErrRespond(w, http.StatusUnauthorized, "Unable to delete other sessions")
		return
	}

	// Step 7: Respond with a success message if all sessions terminated successfully
	respond.ErrRespond(w, http.StatusOK, "All sessions terminated")
}
