package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"
)

const helloMessage = "Hello, World!"

type HelloResponse struct {
	Message string `json:"message,omitempty"`
}

// HelloHandler returns function containing a simple hello world example of an api handler
func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		response := HelloResponse{
			Message: helloMessage,
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Error(ctx, "marshalling response failed", err)
			http.Error(w, "Failed to marshall json response", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}
