package kruskal

import (
	"sort"
)

// Edge is a weighted edge in a graph.
type Edge struct {
	// v1 and v2 are the vertices of the edge.
	v1, v2 int
	// w is the weight of the edge.
	w int
}

// Kruskal returns the minimum spanning tree of the given graph.
func Kruskal(edges []Edge) []Edge {

	// result is the minimum spanning tree.
	var result []Edge

	// sort the edges by weight.
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	// create a new disjoint set.
	uf := newUnionFind(len(edges))
	for _, e := range edges {
		// if the two vertices of the edge are in different sets,
		if !uf.connected(e.v1, e.v2) {
			// unite the two sets.
			uf.union(e.v1, e.v2)
			// add the edge to the result.
			result = append(result, e)
		}
	}

	// return the result.
	return result

}

// unionFind is a union-find data structure.
type unionFind struct {
	// parent is the parent of each node.
	parent []int
	// rank is the rank of each node.
	rank []int
}

// newUnionFind returns a new instance of union-find data structure
func newUnionFind(n int) *unionFind {

	// parents is the parent of each node.
	parents := make([]int, n)

	// ranks is the rank of each node.
	ranks := make([]int, n)

	// initialize the parent and rank of each node.
	for i := 0; i < n; i++ {
		parents[i] = i
		ranks[i] = 0
	}

	// return the new union-find data structure.
	return &unionFind{
		parent: parents,
		rank:   ranks,
	}
}

// find returns the representative of the set that v belongs to.
func (u *unionFind) find(v int) int {

	// if the parent of v is v, return v.
	if u.parent[v] == v {
		return v
	}

	// otherwise, return the root of v.
	return u.find(u.parent[v])
}

// connected returns true if v1 and v2 are connected
func (u *unionFind) connected(v1, v2 int) bool {
	// return true if the two vertices are in the same set.
	return u.find(v1) == u.find(v2)
}

// union performs a union between two disjoint sets
func (u *unionFind) union(v1, v2 int) {

	// find the roots of v1 and v2.
	root1 := u.find(v1)
	root2 := u.find(v2)

	// if the two vertices are in the same set, return.
	if root1 == root2 {
		return
	}

	// if the rank of the root of v1 is less than the rank of the root of v2,
	if u.rank[root1] < u.rank[root2] {
		// set the parent of root1 to root2.
		u.parent[root1] = root2
	} else {
		// set the parent of root2 to root1.
		u.parent[root2] = root1
		// if the ranks are equal, increment the rank of root1.
		if u.rank[root1] == u.rank[root2] {
			u.rank[root1]++
		}
	}
}

// printGraph prints the graph
func printGraph(edges []Edge) {
	for _, edge := range edges {
		println(edge.v1, " -- ", edge.v2, " (", edge.w, ")")
	}
}

// reverseEdge reverses the edge by swapping the vertices
func reverseEdge(edge Edge) Edge {
	return Edge{edge.v2, edge.v1, edge.w}
}

// edgeEqualStrict checks if two edges are equal in strict sense.
func edgeEqualStrict(e1, e2 Edge) bool {
	return e1.v1 == e2.v1 && e1.v2 == e2.v2
}

// edgeEqual checks if two edges are equal.
// an edge is equal to another edge if the two vertices are the same and the weight is the same.
// the comparison also reverses the order of the vertices to check if the edge is the same but in the opposite direction.
// this is not the same as the strict equality of the two edges.
func edgeEqual(e1, e2 Edge) bool {
	return edgeEqualStrict(e1, e2) || edgeEqualStrict(reverseEdge(e1), e2)
}

// graphsEqual returns true iff the two graphs are equal.
func graphsEqual(edges1, edges2 []Edge) bool {
	// fail early if lengths are different
	if len(edges1) != len(edges2) {
		return false
	}
	// edges should be sorted by weight already
	for i := 0; i < len(edges1); i++ {
		if !edgeEqual(edges1[i], edges2[i]) {
			return false
		}
	}
	return true
}
