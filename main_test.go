package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreetHandler_WithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet?name=Alice", nil)
	w := httptest.NewRecorder()
	GreetHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	body := w.Body.String()
	want := "Hello, Alice!"
	if body != want {
		t.Errorf("GreetHandler with name=Alice = %q, want %q", body, want)
	}
}

func TestGreetHandler_DefaultName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()
	GreetHandler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	body := w.Body.String()
	want := "Hello, World!"
	if body != want {
		t.Errorf("GreetHandler with no name = %q, want %q", body, want)
	}
}

func TestGreet(t *testing.T) {
	got := Greet("Alice")
	want := "Hello, Alice!"
	if got != want {
		t.Errorf("Greet(\"Alice\") = %q, want %q", got, want)
	}
}

func TestFarewell(t *testing.T) {
	got := Farewell("Alice")
	want := "Goodbye, Alice!"
	if got != want {
		t.Errorf("Farewell(\"Alice\") = %q, want %q", got, want)
	}
}
