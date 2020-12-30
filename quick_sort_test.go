package gokken

import "testing"
import "reflect"

func TestQuickSort(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
    got := QuickSort(input)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %q want %q", got, want)
    }
}
