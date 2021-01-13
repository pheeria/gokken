package gokken

import "testing"
import "reflect"

func TestBubbleSort(t *testing.T) {
	got := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	BubbleSort(got)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestOptimizedBubbleSort(t *testing.T) {
	got := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	OptimizedBubbleSort(got)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %q want %q", got, want)
    }
}
