package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Doma")
	want := "Hello, Doma!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
