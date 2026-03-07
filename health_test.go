package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler(t *testing.T) {
	startTime = time.Now()
	version = "v1.2.3"

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
}

func TestHealthHandlerUptime(t *testing.T) {
	startTime = time.Now().Add(-2*time.Hour - 15*time.Minute - 30*time.Second)
	version = "dev"

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler().ServeHTTP(rec, req)

	var resp healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// The uptime should contain "h" indicating hours are present.
	if resp.Uptime != "2h15m30s" {
		t.Errorf("expected uptime \"2h15m30s\", got %q", resp.Uptime)
	}
}
