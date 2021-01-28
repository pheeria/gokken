package sort

import "math"

func findMin(array []int) int {
	index := -1
	value := math.MaxInt32

	for i, v := range array {
		if v < value {
			value = v
			index = i
		}
	}

	return index
}

func SelectionSortNewArray(sequence []int) []int {
	result := make([]int, len(sequence))
	length := len(sequence)

	for i := 0; i < length; i++ {
		min := findMin(sequence)
		result[i] = sequence[min]
		sequence = append(sequence[:min], sequence[min+1:]...)
	}

	return result
}

func SelectionSortInPlace(sequence []int) {
	length := len(sequence)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if sequence[j] < sequence[i] {
				swap(sequence, i, j)
			}
		}
	}
}
