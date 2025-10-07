package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	node     string
	distance int
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))

}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

type Graph struct {
	edges map[string]map[string]int
}

func (g *Graph) AddEdge(from, to string, weight int) {
	if g.edges == nil {
		g.edges = make(map[string]map[string]int)
	}
	if g.edges[from] == nil {
		g.edges[from] = make(map[string]int)
	}
	g.edges[from][to] = weight

	g.edges[to][from] = weight
}

func (g *Graph) Dijikstra(start string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	parent := make(map[string]string)
	for node := range g.edges {
		dist[node] = 1<<31 - 1
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		for neighbor, weight := range g.edges[current.node] {
			newDist := current.distance + weight
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				parent[neighbor] = current.node
				heap.Push(pq, Item{node: neighbor, distance: newDist})
			}
		}
	}
	return dist, parent
}

func getPath(parent map[string]string, start, dest string) []string {
	path := []string{}
	for dest != "" {
		path = append([]string{dest}, path...)
		if dest == start {
			break
		}
		dest = parent[dest]
	}
	return path

}

func main() {
	var g Graph
	var edges int

	fmt.Print("Enter number of edges : ")
	fmt.Scan(&edges)

	for i := 0; i < edges; i++ {
		var from, to string
		var weight int
		fmt.Printf("Enter edge %d (example : A B 5): ", i+1)
		fmt.Scan(&from, &to, &weight)
		g.AddEdge(from, to, weight)
	}

	var start, dest string
	fmt.Print("Enter start node : ")
	fmt.Scan(&start)
	fmt.Print("Enter destination node: ")
	fmt.Scan(&dest)

	dist, parent := g.Dijikstra(start)
	path := getPath(parent, start, dest)

	if len(path) == 0 {
		fmt.Println("NO path found between ", start, "add", dest)
	} else {
		fmt.Printf("Shortest path from %s to %s: %v\n", start, dest, path)
		fmt.Printf("Total Distance : %d\n", dist[dest])
	}
}
