package router

import (
	"daemon_backend.bin/component/structure"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func NewRouter(routes structure.Routes) http.Handler {
	router := chi.NewRouter()
	setupSettings(router)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.Get(route.Path, route.Handler)
		case "POST":
			router.Post(route.Path, route.Handler)
		case "PATCH":
			router.Patch(route.Path, route.Handler)
		case "DELETE":
			router.Delete(route.Path, route.Handler)
		case "OPTIONS":
			router.Options(route.Path, route.Handler)
		default:
			log.Printf("Unsupported HTTP method: %s", route.Method)
		}
	}

	return router
}
