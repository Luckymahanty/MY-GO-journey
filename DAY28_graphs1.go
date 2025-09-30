package main

import "fmt"

type Graph struct {
	vertices int
	adjList  map[int][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		adjList:  make(map[int][]int),
	}

}

func (g *Graph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v)

}

func (g *Graph) PrintGraph() {
	for v := 0; v < g.vertices; v++ {
		fmt.Printf("%d -> %v\n", v, g.adjList[v])
	}

}
func main() {
	graph := NewGraph(5)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	fmt.Println("Graph adjacency list:")
	graph.PrintGraph()

}
