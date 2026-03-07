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

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, tt := range tests {
		got := Reverse(tt.input)
		if got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
