package main

import "testing"

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

