package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	BubbleSort(unsorted)
	assertEqual(t, sorted, unsorted)
}

func TestOptimizedBubbleSort(t *testing.T) {
	sorted, unsorted := createTestArrays(10)
	OptimizedBubbleSort(unsorted)
	assertEqual(t, sorted, unsorted)
}
