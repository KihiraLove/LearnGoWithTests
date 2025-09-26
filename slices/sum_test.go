package slices

import (
	"reflect"
	"testing"
)

func assertCorrectSum(t *testing.T, got, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, given)
	}
}

func assertCorrectValueInSlices(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("summing 5 numbers from a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("summing numbers from multiple slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertCorrectValueInSlices(t, got, want)
	})
}
