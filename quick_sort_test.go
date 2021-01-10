package gokken

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	got := QuickSort(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestQuickSortLomuto(t *testing.T) {
	got := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	QuickSortLomuto(got, 0, len(got)-1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestQuickSortHoare(t *testing.T) {
	got := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	QuickSortHoare(got, 0, len(got)-1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
