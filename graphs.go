package gokken

type Node struct {
	Value string
	Edges []Node
}

func FindAktöbe(path []Node, city Node) []Node {
	path = append(path, city)
	var result []Node
	if city.Value == "Aktöbe" {
		return path
	}
	for _, edge := range city.Edges {
		result = FindAktöbe(path, edge)
        if len(result) > 0 {
            return result
        }
	}
	return result
}
