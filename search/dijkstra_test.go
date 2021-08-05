package search

import (
    "testing"
    "reflect"
)

func assertSlices(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func createFirstGraph() Graph {
    start := make(Costs)
    start["a"] = 6
    start["b"] = 2

    a := make(Costs)
    a["finish"] = 1

    b := make(Costs)
    b["a"] = 3
    b["finish"] = 5

    finish := make(Costs)

    result := make(Graph)
    result["start"] = start
    result["a"] = a
    result["b"] = b
    result["finish"] = finish
    return result
}

func createSecondGraph() Graph {
    start := make(Costs)
    start["c"] = 5
    start["f"] = 2

    c := make(Costs)
    c["a"] = 4
    c["e"] = 2

    a := make(Costs)
    a["e"] = 6
    a["finish"] = 3

    f := make(Costs)
    f["c"] = 8
    f["e"] = 7

    e := make(Costs)
    e["finish"] = 1

    finish := make(Costs)

    result := make(Graph)
    result["start"] = start
    result["c"] = c
    result["a"] = a
    result["f"] = f
    result["e"] = e
    result["finish"] = finish
    return result
}

func createThirdGraph() Graph {
    start := make(Costs)
    start["a"] = 10

    a := make(Costs)
    a["c"] = 20

    b := make(Costs)
    b["a"] = 1

    c := make(Costs)
    c["b"] = 1
    c["finish"] = 30

    finish := make(Costs)

    result := make(Graph)
    result["start"] = start
    result["c"] = c
    result["a"] = a
    result["b"] = b
    result["finish"] = finish
    return result
}

func TestDijkstraSearch(t *testing.T) {
	// t.Run("searches the first", func(t *testing.T) {
        // want := []string{"start", "b", "a", "finish"}
        // got := Dijkstra(createFirstGraph(), "start", "finish")
	// 	assertSlices(t, got, want)
	// })
	// t.Run("searches the second", func(t *testing.T) {
        // want := []string{"start", "c", "e", "finish"}
        // got := Dijkstra(createSecondGraph(), "start", "finish")
	// 	assertSlices(t, got, want)
	// })
	t.Run("searches the third", func(t *testing.T) {
        want := []string{"start", "a", "c", "finish"}
        got := Dijkstra(createThirdGraph(), "start", "finish")
		assertSlices(t, got, want)
	})
}
