package sort

import "testing"

func TestSelectionSortInPlace(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	SelectionSortInPlace(unsorted)
	assertEqual(t, sorted, unsorted)
}

func TestSelectionSortNewArray(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	unsorted = SelectionSortNewArray(unsorted)
	assertEqual(t, sorted, unsorted)
}
