package main

import (
	"fmt"
	"math"
)

type Graph struct {
	vertices int
	adjList  map[int][]Edge
}
type Edge struct {
	dest   int
	weight int
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		adjList:  make(map[int][]Edge),
	}

}

func (g *Graph) AddEdge(src, dest, weight int) {
	g.adjList[src] = append(g.adjList[src], Edge{dest, weight})
	g.adjList[dest] = append(g.adjList[dest], Edge{src, weight})

}
func (g *Graph) Dijkstra(start int) []int {
	dist := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	for i := 0; i < g.vertices-1; i++ {
		u := minDistance(dist, visited)
		visited[u] = true

		for _, edge := range g.adjList[u] {
			if !visited[edge.dest] && dist[u]+edge.weight < dist[edge.dest] {
				dist[edge.dest] = dist[u] + edge.weight

			}
		}
	}
	return dist

}

func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1
	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}
	return minIndex

}

func main() {
	graph := NewGraph(6)

	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 5)
	graph.AddEdge(1, 3, 10)
	graph.AddEdge(2, 4, 3)
	graph.AddEdge(4, 3, 4)
	graph.AddEdge(3, 3, 23)

	source := 0
	distances := graph.Dijkstra(source)

	fmt.Printf("Shortest paths from node %d;\n", source)
	for i, d := range distances {
		fmt.Printf("Node %d : %d\n", i, d)
	}

}
