package search

import "testing"

func assert(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestBinary(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7}
	want := 4

	t.Run("searches iteratively", func(t *testing.T) {
		got := IterativeBinarySearch(input, 4)
		assert(t, got, want)
	})

	t.Run("searches recursively", func(t *testing.T) {
		got := RecursiveBinarySearch(input, 4)
		assert(t, got, want)
	})
}

func TestBinaryNotFound(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7}
	want := -1

	t.Run("searches iteratively", func(t *testing.T) {
		got := IterativeBinarySearch(input, 42)
		assert(t, got, want)
	})

	t.Run("searches recursively", func(t *testing.T) {
		got := RecursiveBinarySearch(input, 42)
		assert(t, got, want)
	})
}
