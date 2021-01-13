package gokken

import (
	"fmt"
	"math/rand"
	"sort"
)

func createRandomIntArray(size int) []int {
	var result []int
	rand.Seed(int64(size))

	for size >= len(result) {
        result = append(result, rand.Intn(size * 10))
	}

	return result
}

func sortAsNew(sequence []int) []int {
    result := make([]int, len(sequence))
    copy(result, sequence)
    sort.Ints(result)
    return result
}

func CreateTestArrays(size int) (sorted, unsorted []int) {
    unsorted = createRandomIntArray(size)
    sorted = sortAsNew(unsorted)
    return
}

func Swap(sequence []int, a, b int) {
	tmp := sequence[a]
	sequence[a] = sequence[b]
	sequence[b] = tmp
}

func main() {
    sorted, unsorted := CreateTestArrays(10)

    fmt.Println(sorted)
	fmt.Println(unsorted)
}
