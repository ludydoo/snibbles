package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	graph := NewGraph(4)
	graph.AddEdge(0, 1, 3)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 3)
	graph.AddEdge(1, 3, 3)

	t.Log("\n" + graph.String())

	dijkstra := NewDijkstra(graph)
	dijkstra.Run(0)
	t.Log(dijkstra.PathString(3))

}
