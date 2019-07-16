package api

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error,omitempty"`
}

func writeErrorResponse(rw *responseLogWriter, statusCode int, err error) {
	rw.Header().Set("Content-Type", "application/json")
	if err != nil {
		errResp := &errorResponse{
			Error: err.Error(),
		}
		response, err := json.Marshal(errResp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(statusCode)
		rw.Write(response)
		return
	}
	rw.WriteHeader(statusCode)
}

type withdrawResponse struct {
	Issue string `json:"issue,omitempty"`
}
