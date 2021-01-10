package gokken

func QuickSort(sequence []int) []int {
	if len(sequence) < 2 {
		return sequence
	}

	pivot := sequence[0]
	var less []int
	var greater []int

	for _, v := range sequence[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}

func swap(sequence []int, a, b int) {
	tmp := sequence[a]
	sequence[a] = sequence[b]
	sequence[b] = tmp
}

func partitionLomuto(sequence []int, low, high int) int {
	pivot := sequence[high]
	i := low
	for j := low; j < high; j++ {
		if sequence[j] < pivot {
			swap(sequence, i, j)
			i++
		}
	}
	swap(sequence, i, high)
	return i
}
func QuickSortLomuto(sequence []int, low, high int) {
	if low < high {
		p := partitionLomuto(sequence, low, high)
		QuickSortLomuto(sequence, low, p-1)
		QuickSortLomuto(sequence, p+1, high)
	}
}

func partititionHoare(sequence []int, low, high int) int {
	pivot := (low + high) / 2
	left := low
	right := high
	for {
		for sequence[left] < sequence[pivot] {
			left++
		}
		for sequence[right] > sequence[pivot] {
			right--
		}

		if left >= right {
			return right
		}

		swap(sequence, left, right)
	}
}
func QuickSortHoare(sequence []int, low, high int) {
	if low < high {
		pivot := partititionHoare(sequence, low, high)
		QuickSortHoare(sequence, low, pivot)
		QuickSortHoare(sequence, pivot+1, high)
	}
}
