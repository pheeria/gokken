package gokken

import (
	"reflect"
	"testing"
)

func TestSelectionSortInPlace(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
    SelectionSortInPlace(unsorted)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}

func TestSelectionSortNewArray(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	unsorted = SelectionSortNewArray(unsorted)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}

