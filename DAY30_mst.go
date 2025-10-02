package main

import (
	"fmt"
	"math"
)

type Graph struct {
	vertices int
	adj      [][]int
}

func NewGraph(v int) *Graph {
	adj := make([][]int, v)
	for i := range adj {
		adj[i] = make([]int, v)
	}
	return &Graph{vertices: v, adj: adj}

}

func (g *Graph) AddEdge(u, v, w int) {
	g.adj[u][v] = w
	g.adj[v][u] = w

}

func minKey(key []int, mstSet []bool, V int) int {
	min := math.MaxInt32
	minIndex := -1

	for v := 0; v < V; v++ {
		if !mstSet[v] && key[v] < min {
			min = key[v]
			minIndex = v
		}
	}

	return minIndex

}

func (g *Graph) PrimMST() {
	V := g.vertices
	key := make([]int, V)
	parent := make([]int, V)
	mstSet := make([]bool, V)

	for i := 0; i < V; i++ {
		key[i] = math.MaxInt32
		mstSet[i] = false
	}

	key[0] = 0
	parent[0] = -1

	for count := 0; count < V-1; count++ {
		u := minKey(key, mstSet, V)
		mstSet[u] = true

		for v := 0; v < V; v++ {
			if g.adj[u][v] != 0 && !mstSet[v] && g.adj[u][v] < key[v] {
				parent[v] = u
				key[v] = g.adj[u][v]
			}
		}
	}

	fmt.Println("Edge Weight")
	for i := 1; i < V; i++ {
		fmt.Printf("%d -%d    %d\n", parent[i], i, g.adj[i][parent[i]])
	}

}

func main() {
	g := NewGraph(5)

	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	g.PrimMST()

}
