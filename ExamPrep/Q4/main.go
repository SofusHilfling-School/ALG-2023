package main

import (
	"fmt"
)

func main() {
	g := NewGraph()
	AddEdgeOneWay(g)
	//MakeUndirected(g)

	g.AddVertex(9)

	r := DepthFirstSearch(g, 0)
	MakePath(0, 5, r)
	MakePath(3, 2, r)

	r = BreadthFirstSearch(g, 0)
	MakePath(0, 5, r)
}

func AddEdgeOneWay(g *Graph) {
	g.AddEdge(0, 1, 1)
	g.AddEdge(3, 1, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(5, 2, 1)
	g.AddEdge(5, 4, 1)
	g.AddEdge(4, 6, 1)
	g.AddEdge(6, 8, 1)
	g.AddEdge(1, 7, 1)
	g.AddEdge(2, 6, 1)
	g.AddEdge(7, 8, 1)
	g.AddEdge(7, 2, 1)
}

func MakeUndirected(g *Graph) {
	g.AddEdge(1, 0, 1) //
	g.AddEdge(1, 3, 1) //
	g.AddEdge(5, 3, 1) //
	g.AddEdge(2, 5, 1) //
	g.AddEdge(4, 5, 1) //
	g.AddEdge(6, 4, 1) //
	g.AddEdge(8, 6, 1) //
	g.AddEdge(7, 1, 1) //
	g.AddEdge(6, 2, 1) //
	g.AddEdge(8, 7, 1) //
	g.AddEdge(2, 7, 1) //
}

func DepthFirstSearch(g *Graph, start int) map[int]Edge {
	edgeTo := make(map[int]Edge)
	visited := make(map[int]bool, len(g.adj))

	// keys := maps.Keys(g.adj)

	var DFS func(g *Graph, v int)
	DFS = func(g *Graph, v int) {
		visited[v] = true

		for _, edge := range g.adj[v] {
			if !visited[edge.To] {
				edgeTo[edge.To] = edge
				DFS(g, edge.To)
			}
		}
	}
	DFS(g, start)
	fmt.Println("Visited following nodes:", visited)

	return edgeTo
}

func BreadthFirstSearch(g *Graph, start int) map[int]Edge {
	edgeTo := make(map[int]Edge)
	visited := make(map[int]bool)
	visited[start] = true
	queue := NewQueue(start)

	for !queue.Empty() {
		v := queue.Dequeue()

		for _, edge := range g.adj[v] {
			if !visited[edge.To] {
				edgeTo[edge.To] = edge
				visited[edge.To] = true
				queue.Enqueue(edge.To)
			}
		}
	}

	fmt.Println("Visited following nodes:", visited)
	return edgeTo
}

func MakePath(start, end int, edgeTo map[int]Edge) {
	if curr, ok := edgeTo[end]; ok {
		path := NewStack()
		for curr.To != start {
			path.Push(curr.To)
			curr = edgeTo[curr.From]

		}
		path.Push(start)

		for !path.Empty() {
			v := path.Pop()
			fmt.Print(" -> ", v)
		}
		fmt.Println()
	}
}
