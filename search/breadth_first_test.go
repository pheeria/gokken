package search

import "testing"

func TestBreadthFirstSearch(t *testing.T) {
    outerFirst := Node{ 100, []*Node{} }
    outerSecond := Node{ 200, []*Node{} }
    innerFirst := Node{ 10, []*Node{ &outerFirst } }
    innerSecond := Node{ 20, []*Node{ &outerSecond } }
    root := Node{ 1, []*Node{ &innerFirst, &innerSecond } }

	want := &outerSecond
    got := BreadthFirstSearch(&root, 200)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

