package handler

import (
	"daemon_backend.bin/component/respond"
	"net/http"
)

func Err(w http.ResponseWriter, r *http.Request) {
	respond.ErrRespond(w, 200, "Something went wrong")
}
