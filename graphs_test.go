package gokken

import (
	"fmt"
	"testing"
)

func getCities() Node {
	akt := Node{Value: "Aktöbe"}
	alm := Node{Value: "Almatı", Edges: []Node{akt}}
	ast := Node{Value: "Astana", Edges: []Node{akt, alm}}
	ber := Node{Value: "Berlin", Edges: []Node{ast, alm}}

	return ber
}

func TestFindAktöbe(t *testing.T) {
	ber := getCities()
	var path []Node
	got := FindAktöbe(path, ber)

	fmt.Printf("got %v path %v", got, path)
}
