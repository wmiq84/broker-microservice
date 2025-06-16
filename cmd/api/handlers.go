package main

import (
	"net/http"
)

// creates struct with dummy body values, converts to JSON, sets header and status, then writes to client
// requires response writer b/c send data to client
// read request with request param
// () makes Broker a method on Config
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// plugging in dummy values
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

	// Marshal converts to json
	// out is JSON string, _ is error val
	// so no prefix, and \t for tab space
	// sets Content-Type header to application/json
	// 202 Accepted to client
	// out being the JSON string writes
	// status line - 202 Accepted
	// headers - Content-Type: application/json
	// body - JSON elem
	// write sends JSON to client

	// Originally
	// out, _ := json.MarshalIndent(payload, "", "\t")
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)
}
