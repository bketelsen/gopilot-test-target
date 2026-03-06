package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// startTime records when the server started; set in main().
var startTime time.Time

// Version is set via ldflags at build time, e.g.:
//
//	go build -ldflags "-X main.Version=1.0.0"
var Version = "dev"

// healthResponse is the JSON payload returned by the /health endpoint.
type healthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

// healthHandler returns an http.HandlerFunc that writes the health JSON response.
func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := healthResponse{
			Status:  "ok",
			Uptime:  time.Since(startTime).Round(time.Second).String(),
			Version: Version,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
