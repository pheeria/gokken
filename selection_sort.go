package gokken

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

func SelectionSortNewArray(array []int) []int {
	result := make([]int, len(array))
	length := len(array)

	for i := 0; i < length; i++ {
		min := findMin(array)
		result[i] = array[min]
		array = append(array[:min], array[min+1:]...)
	}

	return result
}

func SelectionSortInPlace(array []int) {
	length := len(array)

	for i := 0; i < length-1; i++ {
		min := array[i]
		for j := i + 1; j < length; j++ {
			if array[j] < min {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}

}
