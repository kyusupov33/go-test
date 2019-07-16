package api

import (
	"log"
	"net/http"
)

type responseLogWriter struct {
	http.ResponseWriter
}

func (r *responseLogWriter) Write(p []byte) {
	n, err := r.ResponseWriter.Write(p)
	if n != len(p) {
		log.Printf("Message `%v` has length=%d, but wrote %d", p, len(p), n)
	}
	if err != nil {
		log.Printf("Message `%v` error on writting: %v", p, err)
	}
}
