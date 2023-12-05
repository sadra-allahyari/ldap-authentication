package server

import (
	"daemon_backend.bin/component/database"
	"daemon_backend.bin/component/extractor"
	"daemon_backend.bin/component/router"
	"daemon_backend.bin/daemon/controller"
	_ "github.com/go-chi/chi/v5"
	"log"
)

// StartServer starts the server
func StartServer() {
	portString := extractor.ExtractStrFromFile("server", "port")

	routes := controller.URLs()
	srv := runServer(router.NewRouter(routes), portString)

	log.Printf("Server starting on port %v", portString)
	database.DbConnector()

	// Run the server indefinitely
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
