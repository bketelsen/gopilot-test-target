package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler(t *testing.T) {
	// Set a known start time so uptime is predictable.
	startTime = time.Now()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler := HealthHandler()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", contentType)
	}

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status \"ok\", got %q", resp.Status)
	}

	if resp.Version == "" {
		t.Error("expected non-empty version")
	}

	if resp.Uptime == "" {
		t.Error("expected non-empty uptime")
	}
}

func TestHealthHandlerVersion(t *testing.T) {
	// Override the version variable to simulate ldflags injection.
	original := version
	version = "1.2.3"
	defer func() { version = original }()

	startTime = time.Now()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler().ServeHTTP(rec, req)

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Version != "1.2.3" {
		t.Errorf("expected version \"1.2.3\", got %q", resp.Version)
	}
}

func TestHealthHandlerUptime(t *testing.T) {
	// Set start time to 2 seconds ago to verify uptime is non-zero.
	startTime = time.Now().Add(-2 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler().ServeHTTP(rec, req)

	var resp HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Parse the uptime duration to verify it's roughly 2 seconds.
	d, err := time.ParseDuration(resp.Uptime)
	if err != nil {
		t.Fatalf("failed to parse uptime duration %q: %v", resp.Uptime, err)
	}

	if d < 1*time.Second || d > 5*time.Second {
		t.Errorf("expected uptime around 2s, got %v", d)
	}
}
