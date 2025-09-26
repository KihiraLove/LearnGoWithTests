package hello

import (
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello", func(t *testing.T) {
		got := Hello("Doma", "English")
		want := "Hello, Doma!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying Hello World when empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Eloide", "Spanish")
		want := "Hola, Eloide!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie!"
		assertCorrectMessage(t, got, want)
	})
}
