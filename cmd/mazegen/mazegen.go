/*
 * MIT License
 *
 * Copyright (c) 2022 Ludovic Cleroux
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package mazegen

import (
	"math/rand"
	"time"
)

func init() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

var maze *Maze

// Maze is the main structure of the maze
type Maze struct {
	// Width of the maze
	Width int
	// Height of the maze
	Height int
	// Maze is the 2D array of cells
	Maze [][]bool
	// es is the remaining edges to be added to the maze
	es intSet
	// us is the disjoint set of vertices
	us *unionSet
}

// reset the maze
func (m *Maze) reset() {
	w := m.Width
	h := m.Height
	// reset the maze
	arr := make([][]bool, h*3)
	for y := 0; y < h; y++ {
		for k := 0; k < 3; k++ {
			arr[y*3+k] = make([]bool, w*3)
		}
		for x := 0; x < w; x++ {
			arr[y*3+1][x*3+1] = true
		}
	}
	m.Maze = arr

	// Given the width and height of the maze, we can calculate the number of edges
	//
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X

	// (3-1)*4 + (4-1)*3 = 17
	edgeCount := (h-1)*w + (w-1)*h
	edgeSet := newIntSet()
	for i := 0; i < edgeCount; i++ {
		edgeSet.add(i)
	}
	m.es = edgeSet
	m.us = newUnionSet(edgeCount)
}

func (m *Maze) Next() bool {
	for len(m.es) > 0 {
		es := m.es
		us := m.us

		// Pick a random edge
		edge := es.random()
		es.remove(edge)

		// Get the two vertices connected by this edge
		v1, v2 := getVertices(edge, m.Width)

		// If they're not in the same set, join them
		if !us.connected(v1, v2) {
			us.union(v1, v2)

			// Remove the wall between the two vertices
			x1, y1 := getCoordinates(v1, m.Width)
			x2, y2 := getCoordinates(v2, m.Width)

			// If the edge is horizontal
			if y1 == y2 {
				if x2 < x1 {
					x1, y1, x2, y2 = x2, y2, x1, y1
				}
				m.Maze[y1*3+1][x1*3+2] = true
				m.Maze[y2*3+1][x2*3] = true
			} else {
				if y2 < y1 {
					x1, y1, x2, y2 = x2, y2, x1, y1
				}
				m.Maze[y1*3+2][x1*3+1] = true
				m.Maze[y2*3][x2*3+1] = true
			}
			return true
		}
	}
	return false
}

// NewMaze creates a new maze with the given width and height
// w is the width of the maze
// h is the height of the maze
func NewMaze(w, h int) *Maze {
	if w < 1 || h < 1 {
		panic("w and h must be greater than 0")
	}
	m := &Maze{
		Width:  w,
		Height: h,
	}
	m.reset()
	return m
}

// getCoordinates returns the x and y coordinates of the given vertex
// v is the vertex number
// w is the width of the maze
func getCoordinates(v, w int) (int, int) {
	return v % w, v / w
}

// getVertices returns the two vertices connected by the given edge
// edge is the edge number
// w is the width of the maze
func getVertices(edge, w int) (int, int) {
	// Given an edge, get the two vertices connected by it
	//
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X

	// We describe the ordering of the edges in the following way:
	// Starting in the top left corner, we scan the edges from left to right, top to bottom.
	// The first edge is the edge connecting the top left node to its right neighbor.
	// Once we reach the right edge, we iterate through the vertical edges
	// The first edge is the edge connecting the top left node to its bottom neighbor.

	// The number of horizontal edges in a row is w-1
	// The number of vertical edges in a row is w
	// The number of rows is h-1
	// The number of edges in a row is w-1 + w

	// Getting the number of edges in a single row
	edgesInRow := 2*(w-1) + 1
	verticesInRow := w

	// Getting the row where the edge is
	rowIdx := edge / edgesInRow

	// Getting the row index
	edgeRowStartIdx := rowIdx * edgesInRow
	vertexRowStartIdx := rowIdx * verticesInRow

	// Getting the edge index in the row
	edgeIndexInRow := edge - edgeRowStartIdx

	horizontalEdgeCount := w - 1

	isHorizontalEdge := edgeIndexInRow < horizontalEdgeCount
	if isHorizontalEdge {
		v1 := vertexRowStartIdx + edgeIndexInRow
		return v1, v1 + 1
	} else {
		v1 := vertexRowStartIdx + edgeIndexInRow - horizontalEdgeCount
		return v1, v1 + verticesInRow
	}
}

// intSet is a set of integers
type intSet map[int]struct{}

// newIntSet creates a new set
func newIntSet() intSet {
	return make(intSet)
}

// random returns a random element from the set
func (s intSet) random() int {
	i := rand.Intn(len(s))
	for k := range s {
		if i == 0 {
			return k
		}
		i--
	}
	return -1
}

// add an item to the set
func (s intSet) add(i int) {
	s[i] = struct{}{}
}

// remove an item from the set
func (s intSet) remove(i int) {
	delete(s, i)
}

// unionSet is a set of disjoint sets
type unionSet struct {
	// parents is a map of the parent of each node
	parents []int
	// ranks is a map of the rank of each node
	ranks []int
}

// newUnionSet creates a new set with the given number of nodes
func newUnionSet(edgeCount int) *unionSet {
	uf := &unionSet{
		parents: make([]int, edgeCount),
		ranks:   make([]int, edgeCount),
	}
	for i := range uf.parents {
		uf.parents[i] = i
		uf.ranks[i] = 0
	}
	return uf
}

// connected returns true if the two nodes are connected
// p and q are the nodes
func (u *unionSet) connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

// find the root of the node
// p is the node to find the root of
func (u *unionSet) find(p int) int {
	if u.parents[p] == p {
		return p
	}
	return u.find(u.parents[p])
}

// union merges the two sets
// p and q are the nodes to merge
func (u *unionSet) union(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)
	if pRoot == qRoot {
		return
	}
	if u.ranks[pRoot] < u.ranks[qRoot] {
		u.parents[pRoot] = qRoot
	} else if u.ranks[pRoot] > u.ranks[qRoot] {
		u.parents[qRoot] = pRoot
	} else {
		u.parents[qRoot] = pRoot
		u.ranks[pRoot]++
	}
}
