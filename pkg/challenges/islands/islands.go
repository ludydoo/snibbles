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

package islands

import "fmt"

// RemoveIslands challenge
//
// A pixel screen with black and white dots is represented as a 2d matrix
// Black dots are represented as 1s, and white dots are represented as 0s.
//
// The goal is to remove as many islands as possible from the screen.
//
// An island is a group of connected black dots.
//
// Only horizontally and vertically connected black dots are considered as
// connected. No diagonal black dots are considered as connected.
//
func RemoveIslands(mat [][]int) [][]int {

	// create disjoint sets from 1s
	// the "root" of the disjoint sets would be the nodes at the edge
	// so the disjoint sets must prefer a parent that is at the edge
	//
	// each vertex is indexed by its position in the matrix
	// the index of the vertex = row * width + column

	// height of the matrix
	height := len(mat)

	// do not consider zero-length matrices
	if height == 0 {
		return mat
	}

	// width of the matrix
	width := len(mat[0])

	// there would be no possibility of islands when any dimension is smaller than 3
	if width < 3 || height < 3 {
		return mat
	}

	// disjoint set
	set := newDset(width, height)

	// iterate through all vertices
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			// vertexId is the index of the vertex
			vertexId := row*width + col
			// if the vertex = 0, continue
			if mat[row][col] == 0 {
				continue
			}
			// if we are not on the last column
			if col < width-1 {
				// this is the value of the right neighbour
				right := mat[row][col+1]
				// if the neighbour is = 1 also, union with the vertex
				if right == 1 {
					set.union(vertexId, vertexId+1)
				}
			}
			// if we are not on the last row
			if row < height-1 {
				// this is the value of the bottom neighbour
				bottom := mat[row+1][col]
				// if the neighbour = 1 also, union with the vertex
				if bottom == 1 {
					set.union(vertexId, vertexId+width)
				}
			}
		}
	}

	// at this point, all neighbouring vertices with index = 1 should be linked together
	// if a vertex is connected to one of the edges, its root will be a node on the side
	// if a vertex is not connected to one of the edges, its root will not be a node on the side

	// iterate through the "middle" nodes only. The "edge" nodes are not considered as islands,
	// so no need to check them
	for row := 1; row < width-1; row++ {
		for col := 1; col < height-1; col++ {
			vIdx := row*width + col
			vertex := mat[row][col]
			if vertex == 0 {
				continue
			}
			parent := set.find(vIdx)
			if !set.isEdge(parent) {
				mat[row][col] = 0
			}
		}
	}

	// print the matrix
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			fmt.Print(mat[row][col], " ")
		}
		fmt.Println()
	}

	return mat
}

// dset represents a disjoint set
type dset struct {
	// parent is the parent of the node
	parent []int
	// rank is the rank of the node
	rank []int
	// w is the width of the matrix
	w int
	// h is the height of the matrix
	h int
}

// newDset creates a new disjoint set
func newDset(w, h int) *dset {
	n := w * h
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &dset{
		parent: parent,
		rank:   rank,
		w:      w,
		h:      h,
	}
}

// connected returns true if the two vertices are connected
func (d *dset) connected(v1, v2 int) bool {
	return d.find(v1) == d.find(v2)
}

// find returns the parent of the node
func (d *dset) find(v int) int {
	if d.parent[v] == v {
		return v
	}
	return d.find(d.parent[v])
}

// union combines the two nodes
func (d *dset) union(v1, v2 int) {
	root1 := d.find(v1)
	root2 := d.find(v2)
	if root1 == root2 {
		return
	}

	// the first vertex is at the edge
	v1IsEdge := d.isEdge(root1)
	// the second vertex is at the edge
	v2IsEdge := d.isEdge(root2)

	if v1IsEdge && !v2IsEdge {
		d.parent[root2] = root1
	} else if v2IsEdge && !v1IsEdge {
		d.parent[root1] = root2
	} else {
		// if none of the vertices is at the edge,
		// the one with the higher rank is the parent
		if d.rank[root1] < d.rank[root2] {
			d.parent[root1] = root2
		} else {
			d.parent[root2] = root1
			if d.rank[root2] == d.rank[root1] {
				d.rank[root1]++
			}
		}
	}
}

func (d *dset) isEdge(root1 int) bool {
	v1Row := root1 / d.w
	v1Col := root1 % d.w
	v1IsEdge := v1Row == 0 || v1Col == 0 || v1Row == d.h-1 || v1Col == d.w-1
	return v1IsEdge
}
