package sort

import (
	"math/rand"
	"reflect"
	gosort "sort"
	"testing"
)

func createRandomIntArray(size int) []int {
	var result []int
	rand.Seed(int64(size))

	for size >= len(result) {
		result = append(result, rand.Intn(size*10))
	}

	return result
}

func sortAsNew(sequence []int) []int {
	result := make([]int, len(sequence))
	copy(result, sequence)
	gosort.Ints(result)
	return result
}

func createTestArrays(size int) (sorted, unsorted []int) {
	unsorted = createRandomIntArray(size)
	sorted = sortAsNew(unsorted)
	return
}

func swap(sequence []int, a, b int) {
	tmp := sequence[a]
	sequence[a] = sequence[b]
	sequence[b] = tmp
}

func assertEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
