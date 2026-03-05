package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler_StatusOK(t *testing.T) {
	startTime = time.Now().Add(-5 * time.Minute)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("healthHandler returned status %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestHealthHandler_ContentTypeJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("healthHandler Content-Type = %q, want %q", ct, "application/json")
	}
}

func TestHealthHandler_ResponseBody(t *testing.T) {
	startTime = time.Now().Add(-10 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("Status = %q, want %q", resp.Status, "ok")
	}
	if resp.Version == "" {
		t.Error("Version should not be empty")
	}
	if resp.Uptime == "" {
		t.Error("Uptime should not be empty")
	}
	if resp.Dependencies == nil {
		t.Error("Dependencies should not be nil")
	}
}

func TestHealthHandler_UptimeIncreases(t *testing.T) {
	startTime = time.Now().Add(-30 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Uptime == "" {
		t.Error("Uptime should not be empty")
	}
}

func TestHealthHandler_DependencyStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	for name, dep := range resp.Dependencies {
		if dep.Status == "" {
			t.Errorf("dependency %q has empty status", name)
		}
	}
}
