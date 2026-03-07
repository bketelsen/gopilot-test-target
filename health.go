package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// version is set via ldflags at build time.
var version = "dev"

// startTime records when the server started.
var startTime time.Time

func init() {
	startTime = time.Now()
}

// healthResponse represents the JSON payload returned by the health endpoint.
type healthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

// healthHandler returns an http.HandlerFunc that reports server health.
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
