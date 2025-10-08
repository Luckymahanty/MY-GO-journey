package main

import (
	"fmt"
	"math"
)

type Edge struct {
	from, to string
	weight   int
}

type Graph struct {
	edges []Edge
	nodes map[string]bool
}

func (g *Graph) AddEdge(from, to string, weight int) {
	g.edges = append(g.edges, Edge{from, to, weight})
	if g.nodes == nil {
		g.nodes = make(map[string]bool)
	}
	g.nodes[from] = true
	g.nodes[to] = true
}

func (g *Graph) BellmanFord(start string) (map[string]int, map[string]string, bool) {
	dist := make(map[string]int)
	parent := make(map[string]string)

	for node := range g.nodes {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	for i := 0; i < len(g.nodes)-1; i++ {
		for _, edge := range g.edges {
			if dist[edge.from] != math.MaxInt32 && dist[edge.from]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[edge.from] + edge.weight
				parent[edge.to] = edge.from
			}
		}
	}

	for _, edge := range g.edges {
		if dist[edge.from] != math.MaxInt32 && dist[edge.from]+edge.weight < dist[edge.to] {
			return dist, parent, true // negative cycle detected
		}
	}

	return dist, parent, false
}

func getPath(parent map[string]string, start, dest string) []string {
	path := []string{}
	curr := dest
	for curr != "" {
		path = append([]string{curr}, path...)
		curr = parent[curr]
		if curr == start {
			path = append([]string{start}, path...)
			break
		}
	}
	return path
}

func main() {
	var g Graph
	var edges int

	fmt.Print("Enter number of edges: ")
	fmt.Scan(&edges)

	for i := 0; i < edges; i++ {
		var from, to string
		var weight int
		fmt.Printf("Enter edge %d (example: A B -5): ", i+1)
		fmt.Scan(&from, &to, &weight)
		g.AddEdge(from, to, weight)
	}

	var start, dest string
	fmt.Print("Enter start node: ")
	fmt.Scan(&start)
	fmt.Print("Enter destination node: ")
	fmt.Scan(&dest)

	dist, parent, hasNegativeCycle := g.BellmanFord(start)

	if hasNegativeCycle {
		fmt.Println("❌ Negative weight cycle detected. Shortest path not possible.")
		return
	}

	path := getPath(parent, start, dest)

	if len(path) == 0 {
		fmt.Println("❌ No path found between", start, "and", dest)
	} else {
		fmt.Printf("✅ Shortest Path from %s to %s: %v\n", start, dest, path)
		fmt.Printf("🧭 Total Distance: %d\n", dist[dest])
	}
}
