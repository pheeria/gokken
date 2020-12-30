package gokken

import "testing"
import "reflect"

func TestSelectionSortInPlace(t *testing.T) {
	got := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}
	SelectionSortInPlace(got)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestSelectionSortNewArray(t *testing.T) {
	got := SelectionSortNewArray([]int{5, 4, 3, 2, 1})
	want := []int{1, 2, 3, 4, 5}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %q want %q", got, want)
    }
}

