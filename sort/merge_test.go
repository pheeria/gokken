package sort

import "testing"

func TestMergeSort(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	unsorted = MergeSort(unsorted)
	assertEqual(t, sorted, unsorted)
}
