package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// startTime records when the server started.
var startTime time.Time

// version is set via ldflags at build time (e.g. -ldflags "-X main.version=v1.0.0").
var version = "dev"

// healthResponse is the JSON structure returned by the /health endpoint.
type healthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

// healthHandler returns an http.HandlerFunc that serves the /health endpoint.
func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := healthResponse{
			Status:  "ok",
			Uptime:  time.Since(startTime).Round(time.Second).String(),
			Version: version,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
