package handler

import (
	"daemon_backend.bin/component/respond"
	"net/http"
)

func Readiness(w http.ResponseWriter, r *http.Request) {
	respond.JsonRespond(w, 200, struct{}{})
}
