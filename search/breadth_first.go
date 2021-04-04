package search

type Node struct {
    value int
    edges []*Node
}

func BreadthFirstSearch(root *Node, target int) *Node {
    queue := make([]*Node, 0)
    queue = append(queue, root)
    seen := make(map[int]bool)

    for len(queue) > 0 {
        current := queue[0]

        if current.value == target {
            return current
        }

        if _, exists := seen[current.value]; !exists {
            seen[current.value] = true
        }

        for _, edge := range current.edges {
            if _, exists := seen[edge.value]; !exists {
                queue = append(queue, edge)
            }
        }

        queue = queue[1:]
    }

    return nil
}
