package main

import (
	"encoding/json"
	"net/http"
)

// struct tags that rename field i.e. Error to error
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// requires response writer b/c send data to client
// read request with request param
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// plugging in dummy values
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	// Marshal converts to json
	// out is JSON string, _ is error val
	// so no prefix, and \t for tab space
	out, _ := json.MarshalIndent(payload, "", "\t")
	// sets Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// 202 Accepted to client
	w.WriteHeader(http.StatusAccepted)
	// out being the JSON string writes
	// status line - 202 Accepted
	// headers - Content-Type: application/json
	// body - JSON elem
	w.Write(out)
}
