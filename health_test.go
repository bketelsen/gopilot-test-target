package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler_StatusOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %q", ct)
	}
}

func TestHealthHandler_ResponseFields(t *testing.T) {
	// Set known values for deterministic testing.
	startTime = time.Now().Add(-2 * time.Hour)
	version = "v1.2.3"

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status \"ok\", got %q", resp.Status)
	}

	if resp.Version != "v1.2.3" {
		t.Errorf("expected version \"v1.2.3\", got %q", resp.Version)
	}

	if resp.Uptime == "" {
		t.Error("expected non-empty uptime")
	}

	// Uptime should contain "h" since we set startTime 2 hours ago.
	if resp.Uptime[:2] != "2h" {
		t.Errorf("expected uptime to start with \"2h\", got %q", resp.Uptime)
	}
}

func TestHealthHandler_DefaultVersion(t *testing.T) {
	version = "dev"
	startTime = time.Now()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Version != "dev" {
		t.Errorf("expected version \"dev\", got %q", resp.Version)
	}
}
