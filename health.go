package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// version is set via ldflags at build time, e.g.:
//
//	go build -ldflags "-X main.version=1.0.0"
var version = "dev"

// startTime records when the server started.
var startTime time.Time

func init() {
	startTime = time.Now()
}

// HealthResponse represents the JSON payload returned by the /health endpoint.
type HealthResponse struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime"`
	Version string `json:"version"`
}

// HealthHandler returns an http.HandlerFunc that responds with server health information.
func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := HealthResponse{
			Status:  "ok",
			Uptime:  time.Since(startTime).Round(time.Second).String(),
			Version: version,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
