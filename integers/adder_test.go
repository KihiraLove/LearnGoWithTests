package integers

import (
	"testing"
)

func assertCorrectValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectValue(t, sum, expected)
	})
	t.Run("add two numbers when one is negative", func(t *testing.T) {
		sum := Add(2, -3)
		expected := -1

		assertCorrectValue(t, sum, expected)
	})
}
