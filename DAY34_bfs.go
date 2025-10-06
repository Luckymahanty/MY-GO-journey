package main

import (
	"fmt"

	"bufio"
	"os"
	"strings"
)

type Graph struct {
	vertices map[string][]string
}

func (g *Graph) AddEdge(v1, v2 string) {
	if g.vertices == nil {
		g.vertices = make(map[string][]string)
	}
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)

}
func (g *Graph) ShottestPath(src, dest string) []string {
	queue := []string{src}
	visited := make(map[string]bool)
	parent := make(map[string]string)

	visited[src] = true
	parent[src] = ""

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == dest {
			break
		}

		for _, neighbor := range g.vertices[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
				parent[neighbor] = node
			}
		}
	}

	path := []string{}
	curr := dest
	for curr != "" {
		path = append([]string{curr}, path...)
		curr = parent[curr]
	}
	return path
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	g := Graph{}

	var edges int
	fmt.Print("Enter number of edgrs: ")
	fmt.Scan(&edges)

	for i := 0; i < edges; i++ {
		fmt.Printf("Enter edge %d (example: A B): ", i+1)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		nodes := strings.Split(text, " ")
		if len(nodes) == 2 {
			g.AddEdge(nodes[0], nodes[1])
		} else {
			fmt.Println("Invalid input, enter two nodes sepated by space.")
			i--
		}
	}

	fmt.Print("Enter start node: ")
	src, _ := reader.ReadString('\n')
	src = strings.TrimSpace(src)

	fmt.Print("Enter destination node : ")
	dest, _ := reader.ReadString('\n')
	dest = strings.TrimSpace(dest)

	path := g.ShottestPath(src, dest)

	if len(path) == 0 {
		fmt.Println(" NO path found between", src, "and", dest)
	} else {
		fmt.Println("Shortest path from ", src, "to", dest, "is:", strings.Join(path, " ->"))
	}

}
