package gokken

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
    sorted, unsorted := CreateTestArrays(10)
	unsorted = MergeSort(unsorted)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("got %v want %v", unsorted, sorted)
	}
}

