package main

type Edge struct {
	To     int
	From   int
	Weight int
}

type Graph struct {
	adj map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]Edge),
	}
}

func (g *Graph) AddVertex(v int) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = make([]Edge, 0)
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)
	e := Edge{
		To:     to,
		From:   from,
		Weight: weight,
	}
	g.adj[from] = append(g.adj[from], e)
}
