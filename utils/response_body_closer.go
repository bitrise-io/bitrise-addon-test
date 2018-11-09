package utils

import (
	"log"
	"net/http"
)

// ResponseBodyCloseWithErrorLog ...
func ResponseBodyCloseWithErrorLog(r *http.Response) {
	err := r.Body.Close()
	if err != nil {
		log.Printf(" [!] Exception: ResponseBodyCloseWithErrorLog: %+v", err)
	}
}
