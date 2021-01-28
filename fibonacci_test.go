package gokken

import (
	"reflect"
	"testing"
)

var want = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

func TestIterativeFibonacci(t *testing.T) {
	got := IterativeFibonacci(10)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRecursiveFibonacci(t *testing.T) {
	got := RecursiveFibonacci(10)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestClosureFibonacci(t *testing.T) {
	got := ClosureFibonacci(10)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
