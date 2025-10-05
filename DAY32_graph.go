package main

import (
	"bufio"
	"fmt"
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
func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	fmt.Println("BFS Traversal: ")
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			fmt.Println(node, " ")
			visited[node] = true
			for _, neighbor := range g.vertices[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	fmt.Println()

}

func (g *Graph) DFS(start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	fmt.Println(start, " ")
	visited[start] = true

	for _, neighbor := range g.vertices[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	g := Graph{}

	var edges int
	fmt.Print("Enter number of edges: ")
	fmt.Scan(&edges)

	for i := 0; i < edges; i++ {
		fmt.Printf("Enter edge %d (example : A B): ", i+1)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		nodes := strings.Split(text, " ")
		if len(nodes) == 2 {
			g.AddEdge(nodes[0], nodes[1])
		} else {
			fmt.Println(" Invalid input, please enter two nodes separated by space :")
			i--
		}
	}
	fmt.Print("Enter staring node: ")
	startNode, _ := reader.ReadString('\n')
	startNode = strings.TrimSpace(startNode)

	fmt.Print("Choose traversal (BFS/DFS): ")
	traversalType, _ := reader.ReadString('\n')
	traversalType = strings.ToUpper(strings.TrimSpace(traversalType))

	if traversalType == "BFS" {
		fmt.Println("DFS Traversal : ")
		visited := make(map[string]bool)
		g.DFS(startNode, visited)
		fmt.Println()
	} else {
		fmt.Println("Invalid traversal type!")
	}

}
