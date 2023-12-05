package structure

import (
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route
