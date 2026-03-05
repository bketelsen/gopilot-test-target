package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Version is the application version, set at build time.
var Version = "dev"

// startTime records when the service started, used to compute uptime.
var startTime = time.Now()

// DependencyStatus represents the health status of a single dependency.
type DependencyStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// HealthResponse is the JSON body returned by the /health endpoint.
type HealthResponse struct {
	Status       string                      `json:"status"`
	Version      string                      `json:"version"`
	Uptime       string                      `json:"uptime"`
	Dependencies map[string]DependencyStatus `json:"dependencies"`
}

// checkDependencies returns a map of dependency names to their health status.
// Add real dependency checks here as the service grows.
func checkDependencies() map[string]DependencyStatus {
	return map[string]DependencyStatus{
		"self": {Status: "ok"},
	}
}

// healthHandler handles GET /health and returns service health as JSON.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	deps := checkDependencies()

	overallStatus := "ok"
	for _, d := range deps {
		if d.Status != "ok" {
			overallStatus = "degraded"
			break
		}
	}

	uptime := time.Since(startTime).Truncate(time.Second)

	resp := HealthResponse{
		Status:       overallStatus,
		Version:      Version,
		Uptime:       fmt.Sprintf("%s", uptime),
		Dependencies: deps,
	}

	w.Header().Set("Content-Type", "application/json")
	if overallStatus != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(resp)
}
