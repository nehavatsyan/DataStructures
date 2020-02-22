package graph

type wgraph struct {
	n     int
	edges [][]edge
}
type edge struct {
	from   int
	to     int
	weight int
}

type opsWgraph interface {
	createWgraph()
	addEdge()
}

func (g wgraph) createWgraph(n int) *wgraph {
	g.n = n
	return &g
}

func (g *wgraph) addEdge(u, v, w int) *wgraph {
	g.edges[u] = append(g.edges[u], edge{from: u, to: v, weight: w})
	return g
}
