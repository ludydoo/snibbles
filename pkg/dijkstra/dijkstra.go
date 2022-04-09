package dijkstra

import (
	"strconv"
	"strings"
)

// Edge is a weighted edge in a graph.
type Edge struct {
	// To is the destination vertex.
	To int
	// Cost is the cost of the edge.
	Cost int
}

// String returns a string representation of the edge.
func (e Edge) String() string {
	return "(" + strconv.Itoa(e.To) + "," + strconv.Itoa(e.Cost) + ")"
}

// Graph is a graph represented as an adjacency list.
type Graph struct {
	// Edges is a map of vertices to their edges.
	Edges [][]Edge
}

// NewGraph returns a new graph with the given number of vertices.
func NewGraph(n int) *Graph {
	g := &Graph{}
	g.Edges = make([][]Edge, n)
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
	}
	return g
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Cost: cost})
}

// String returns a string representation of the graph.
func (g *Graph) String() string {
	var sb strings.Builder
	for i, edges := range g.Edges {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString(": ")
		for _, edge := range edges {
			sb.WriteString(strconv.Itoa(edge.To))
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(edge.Cost))
			sb.WriteString(") ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// Dijkstra algorithm
type Dijkstra struct {
	// Graph is the graph to run Dijkstra's on.
	Graph *Graph
	// Dist is the distance from the source to each vertex.
	Dist []int
	// Prev is the previous vertex in the shortest path from the source to
	Prev []int
	// Queued is a list of vertices that have been queued for processing.
	Queued []bool
}

// NewDijkstra returns a new Dijkstra instance.
func NewDijkstra(g *Graph) *Dijkstra {
	d := &Dijkstra{
		Graph:  g,
		Dist:   make([]int, len(g.Edges)),
		Prev:   make([]int, len(g.Edges)),
		Queued: make([]bool, len(g.Edges)),
	}
	for i := 0; i < len(g.Edges); i++ {
		d.Dist[i] = -1
		d.Prev[i] = -1
		d.Queued[i] = false
	}
	return d
}

// Run Dijkstra's algorithm on the graph g, starting at start.
func (d *Dijkstra) Run(start int) {
	d.Dist[start] = 0
	d.Queued[start] = true
	for len(d.Queued) > 0 {
		min := -1
		minDist := -1
		for i, q := range d.Queued {
			if q && (min == -1 || d.Dist[i] < minDist) {
				min = i
				minDist = d.Dist[i]
			}
		}
		if min == -1 {
			break
		}
		d.Queued[min] = false
		for _, edge := range d.Graph.Edges[min] {
			if d.Dist[edge.To] == -1 || d.Dist[edge.To] > d.Dist[min]+edge.Cost {
				d.Dist[edge.To] = d.Dist[min] + edge.Cost
				d.Prev[edge.To] = min
				d.Queued[edge.To] = true
			}
		}
	}
}

// Path returns the path from start to end.
func (d *Dijkstra) Path(t int) []int {
	path := make([]int, 0)
	for t != -1 {
		path = append(path, t)
		t = d.Prev[t]
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

// PathCost returns the cost of the path from start to end.
func (d *Dijkstra) PathCost(s, t int) int {
	return d.Dist[t]
}

// PathString returns the path from start to end as a string.
func (d *Dijkstra) PathString(t int) string {
	path := d.Path(t)
	var pathStrs []string
	for _, i := range path {
		pathStrs = append(pathStrs, strconv.Itoa(i))
	}
	return "->" + strings.Join(pathStrs, "->")
}

// PathStringCost returns the cost of the path from start to end as a string.
func (d *Dijkstra) PathStringCost(s, t int) string {
	return strconv.Itoa(d.PathCost(s, t))
}
