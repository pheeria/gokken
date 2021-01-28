package sort

import "testing"

func TestQuickSort(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	unsorted = QuickSort(unsorted)
	assertEqual(t, sorted, unsorted)
}

func TestQuickSortLomuto(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	QuickSortLomuto(unsorted, 0, len(unsorted)-1)
	assertEqual(t, sorted, unsorted)
}

func TestQuickSortHoare(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	QuickSortHoare(unsorted, 0, len(unsorted)-1)
	assertEqual(t, sorted, unsorted)
}
