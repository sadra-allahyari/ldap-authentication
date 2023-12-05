package controller

import (
	"daemon_backend.bin/component/handler"
	"daemon_backend.bin/component/middleware"
	"daemon_backend.bin/component/structure"
	"net/http"
)

const apiPrefix = "/api/v2/"

func URLs() structure.Routes {
	var routes structure.Routes

	authMiddleware := middleware.AuthMiddleware

	routes = structure.Routes{
		{Method: "GET", Path: apiPrefix + "healthz", Handler: handler.Readiness},
		{Method: "GET", Path: apiPrefix + "test", Handler: handler.Err},
		{Method: "POST", Path: apiPrefix + "user/login", Handler: handler.LoginHandler},
		{Method: "DELETE", Path: apiPrefix + "user/terminate", Handler: authMiddleware(http.HandlerFunc(handler.TerminateHandler)).ServeHTTP},
	}
	return routes
}
