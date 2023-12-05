package respond

import (
	"log"
	"net/http"
)

func ErrRespond(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println(msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	JsonRespond(w, code, errResponse{
		Error: msg,
	})
}
