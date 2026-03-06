package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler_StatusOK(t *testing.T) {
	startTime = time.Now()
	Version = "1.2.3"

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	ct := w.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	var resp healthResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status %q, got %q", "ok", resp.Status)
	}
	if resp.Version != "1.2.3" {
		t.Errorf("expected version %q, got %q", "1.2.3", resp.Version)
	}
	if resp.Uptime == "" {
		t.Error("expected non-empty uptime")
	}
}

func TestHealthHandler_UptimeAdvances(t *testing.T) {
	startTime = time.Now().Add(-2*time.Hour - 15*time.Minute - 30*time.Second)
	Version = "test"

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler().ServeHTTP(w, req)

	var resp healthResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Parse the uptime duration to verify it's roughly correct
	d, err := time.ParseDuration(resp.Uptime)
	if err != nil {
		t.Fatalf("uptime %q is not a valid duration: %v", resp.Uptime, err)
	}

	if d < 2*time.Hour+15*time.Minute {
		t.Errorf("expected uptime >= 2h15m, got %s", resp.Uptime)
	}
}

func TestHealthHandler_DefaultVersion(t *testing.T) {
	startTime = time.Now()
	Version = "dev"

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler().ServeHTTP(w, req)

	var resp healthResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Version != "dev" {
		t.Errorf("expected default version %q, got %q", "dev", resp.Version)
	}
}
