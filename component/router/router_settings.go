package router

import (
	"daemon_backend.bin/component/extractor"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func setupSettings(router *chi.Mux) {
	allowedOrigins := extractor.ExtractStrFromFile("router", "allowedOrigins")
	allowedMethods := extractor.ExtractStrFromFile("router", "allowedMethods")
	allowedHeaders := extractor.ExtractStrFromFile("router", "allowedHeaders")
	exposedHeaders := extractor.ExtractStrFromFile("router", "exposedHeaders")
	allowCredentials := extractor.ExtractBoolFromFile("router", "allowCredentials")
	maxAge := extractor.ExtractIntFromFile("router", "maxAge")

	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{allowedOrigins},
		AllowedMethods:   []string{allowedMethods},
		AllowedHeaders:   []string{allowedHeaders},
		ExposedHeaders:   []string{exposedHeaders},
		AllowCredentials: allowCredentials,
		MaxAge:           maxAge,
	}))
}
