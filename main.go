package main

import (
	"daemon_backend.bin/component/server"
	"daemon_backend.bin/component/update"
	_ "github.com/lib/pq"
)

func main() {

	// Run the update.sh file
	if err := update.RunUpdateFile(); err != nil {
		panic(err)
	}

	// Run server
	server.StartServer()
}
