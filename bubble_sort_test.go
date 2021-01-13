package gokken

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	BubbleSort(unsorted)
	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
    }
}

func TestOptimizedBubbleSort(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	OptimizedBubbleSort(unsorted)
	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
    }
}
