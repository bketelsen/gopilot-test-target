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

	healthHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", contentType)
	}

	var resp healthResponse
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
	original := version
	defer func() { version = original }()

	version = "v1.2.3"
	startTime = time.Now()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Version != "v1.2.3" {
		t.Errorf("expected version \"v1.2.3\", got %q", resp.Version)
	}
}
