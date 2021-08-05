package search

import (
	"fmt"
	"math"
)

type Costs map[string]int
type Parents map[string]string
type Graph map[string]Costs

func findCheapest(seen map[string]bool, costs Costs) string {
    result := ""
    tmp := math.MaxInt32

    for k, v := range costs {
        if _, exists := seen[k]; !exists && v < tmp {
            tmp = v
            result = k
        }
    }

    return result
}

func Dijkstra(graph Graph, start, finish string) []string {
    costs := make(Costs)
    parents := make(Parents)
    seen := make(map[string]bool)

    for node, dict := range graph {
        for k, v := range dict {
            if node == start {
                fmt.Printf("%s -> %d", k, v)
                costs[k] = v
            } else {
                costs[k] = math.MaxInt32
            }
            parents[k] = node
        }
    }
    parents[finish] = ""
    costs[finish] = math.MaxInt32

    node := findCheapest(seen, costs)
    for node != "" {
        cost := costs[node]
        neighbors := graph[node]

        for k, v := range neighbors {
            newCost := cost + v
            if costs[k] > newCost {
                costs[k] = newCost
                parents[k] = node
            }
        }
        seen[node] = true
        node = findCheapest(seen, costs)
    }

    fmt.Println(costs)

    result := []string{ finish }
    current := finish
    i := 0
    for current != start {
        i++
        if i > 20 { break }
        current = parents[current]
        result = append(result, current)
    }
    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }
    fmt.Println(result)
    return result
}
