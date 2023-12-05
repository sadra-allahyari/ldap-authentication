package server

import "net/http"

func runServer(handler http.Handler, port string) *http.Server {
	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}
	return srv
}
