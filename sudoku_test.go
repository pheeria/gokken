package gokken

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
    got := [][]int{
        {9,7,6,1,0,0,0,0,5},
        {0,0,5,0,9,0,2,0,1},
        {8,0,0,0,4,0,0,0,0},
        {0,0,0,0,8,0,0,0,0},
        {0,0,0,7,0,0,0,0,0},
        {0,0,7,0,2,6,0,0,9},
        {2,0,0,3,0,0,0,0,6},
        {0,0,0,2,0,0,9,0,0},
        {0,0,1,9,0,4,5,7,0},
    }
    want := [][]int{
        {9,7,6,1,3,2,8,4,5},
        {4,3,5,8,9,7,2,6,1},
        {8,1,2,6,4,5,7,9,3},
        {6,2,3,4,8,9,1,5,7},
        {5,9,8,7,1,3,6,2,4},
        {1,4,7,5,2,6,3,8,9},
        {2,5,9,3,7,8,4,1,6},
        {7,6,4,2,5,1,9,3,8},
        {3,8,1,9,6,4,5,7,2},
    }
    Solve(got)

    if !reflect.DeepEqual(got, want) {
        t.Error("Got:\n")
        for i := 0; i < 9; i++ {
            t.Errorf("%v", got[i])
        }
        t.Error("Wanted:\n")
        for i := 0; i < 9; i++ {
            t.Errorf("%v", want[i])
        }
    }
}

