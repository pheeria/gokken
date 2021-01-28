package sort

import "testing"

var _, small = createTestArrays(10)
var _, medium = createTestArrays(100)
var _, large = createTestArrays(1000)

func benchmarkBubble(array []int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		BubbleSort(array)
	}
}
func BenchmarkBubbleSmall(b *testing.B)  { benchmarkBubble(small, b) }
func BenchmarkBubbleMedium(b *testing.B) { benchmarkBubble(medium, b) }
func BenchmarkBubbleLarge(b *testing.B)  { benchmarkBubble(large, b) }

func benchmarkSelection(array []int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SelectionSortInPlace(array)
	}
}
func BenchmarkSelectionSmall(b *testing.B)  { benchmarkSelection(small, b) }
func BenchmarkSelectionMedium(b *testing.B) { benchmarkSelection(medium, b) }
func BenchmarkSelectionLarge(b *testing.B)  { benchmarkSelection(large, b) }

func benchmarkMerge(array []int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MergeSort(array)
	}
}
func BenchmarkMergeSmall(b *testing.B)  { benchmarkMerge(small, b) }
func BenchmarkMergeMedium(b *testing.B) { benchmarkMerge(medium, b) }
func BenchmarkMergeLarge(b *testing.B)  { benchmarkMerge(large, b) }

func benchmarkQuick(array []int, b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		QuickSortHoare(array, 0, len(array)-1)
	}
}
func BenchmarkQuickSmall(b *testing.B)  { benchmarkQuick(small, b) }
func BenchmarkQuickMedium(b *testing.B) { benchmarkQuick(medium, b) }
func BenchmarkQuickLarge(b *testing.B)  { benchmarkQuick(large, b) }
