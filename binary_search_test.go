package gokken

import "testing"

func TestBinarySearch(t *testing.T) {
    input := []int { 0, 1, 2, 3, 4, 5, 6, 7}
    want := 4

    assert := func(t *testing.T, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    }

    t.Run("searches iteratively", func(t *testing.T) {
        got := IterativeBinarySearch(input, 4)
        assert(t, got, want)
    })

    t.Run("searches recursively", func(t *testing.T) {
        got := RecursiveBinarySearch(input, 0, len(input) - 1, 4)
        assert(t, got, want)
    })
}

func TestBinarySearchNotFound(t *testing.T) {
    input := []int { 0, 1, 2, 3, 4, 5, 6, 7}
    want := -1

    assert := func(t *testing.T, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    }

    t.Run("searches iteratively", func(t *testing.T) {
        got := IterativeBinarySearch(input, 42)
        assert(t, got, want)
    })

    t.Run("searches recursively", func(t *testing.T) {
        got := RecursiveBinarySearch(input, 0, len(input) - 1, 42)
        assert(t, got, want)
    })
}


