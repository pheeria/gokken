package gokken

import (
	"math"
	"testing"
)

func TestRecursiveSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	got := RecursiveSum(input)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestRecursiveCount(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	got := RecursiveCount(input)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestRecursiveMax(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	got := RecursiveMax(input, math.MinInt32)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
