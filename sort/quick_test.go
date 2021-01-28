package gokken

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	unsorted = QuickSort(unsorted)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}

func TestQuickSortLomuto(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
    QuickSortLomuto(unsorted, 0, len(unsorted) - 1)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}

func TestQuickSortHoare(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	QuickSortHoare(unsorted, 0, len(unsorted) - 1)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}
