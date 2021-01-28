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

func partitionLomuto(sequence []int, low, high int) int {
	pivot := sequence[high]
	i := low
	for j := low; j < high; j++ {
		if sequence[j] < pivot {
			Swap(sequence, i, j)
			i++
		}
	}
	Swap(sequence, i, high)
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
	pivot := sequence[(low + high) / 2]
	left := low
	right := high
	for left <= right {
		for sequence[left] < pivot {
			left++
		}
		for sequence[right] > pivot {
			right--
		}

		if left <= right {
            Swap(sequence, left, right)
            left++
            right--
		}
	}

    return left
}
func QuickSortHoare(sequence []int, low, high int) {
	if low < high {
		pivot := partititionHoare(sequence, low, high)
		QuickSortHoare(sequence, low, pivot - 1)
		QuickSortHoare(sequence, pivot, high)
	}
}
