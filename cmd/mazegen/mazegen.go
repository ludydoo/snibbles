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
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var maze [][]bool

func createNewMaze(callback func()) {
	maze = NewMaze(60, 50, callback)
}

func Run() error {

	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()
	evQueue := make(chan termbox.Event)
	go func() {
		for {
			evQueue <- termbox.PollEvent()
		}
	}()
	createNewMaze(func() {

	})
	draw()
loop:
	for {
		select {
		case ev := <-evQueue:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					break loop
				} else {
					createNewMaze(func() {
						draw()
					})
				}
			}
		default:
			draw()
			// time.Sleep(time.Millisecond)
		}
	}

	return nil
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] {
				termbox.SetBg(x*2, y, termbox.ColorWhite)
				termbox.SetBg(x*2+1, y, termbox.ColorWhite)
			} else {
				termbox.SetBg(x*2, y, termbox.ColorBlack)
				termbox.SetBg(x*2+1, y, termbox.ColorBlack)
			}
		}
	}
	termbox.Flush()
}

func NewMaze(w, h int, callback func()) [][]bool {

	// Given the width and height of the maze, we can calculate the number of edges
	//
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X
	//  |   |   |   |
	//  X - X - X - X

	// (3-1)*4 + (4-1)*3 = 17
	edgesInMaze := (h-1)*w + (w-1)*h

	edgeSet := newIntSet()
	for i := 0; i < edgesInMaze; i++ {
		edgeSet.add(i)
	}

	// make the maze
	maze = make([][]bool, h*3)
	for y := 0; y < h; y++ {
		for k := 0; k < 3; k++ {
			maze[y*3+k] = make([]bool, w*3)
		}
		for x := 0; x < w; x++ {
			maze[y*3+1][x*3+1] = true
		}
	}

	uf := newUnionSet(edgesInMaze)
	for len(edgeSet) > 0 {
		// Pick a random edge
		edge := edgeSet.random()
		edgeSet.remove(edge)

		// Get the two vertices connected by this edge
		v1, v2 := getVertices(edge, w)

		// If they're not in the same set, join them
		if !uf.connected(v1, v2) {
			uf.union(v1, v2)

			// Remove the wall between the two vertices
			x1, y1 := getCoordinates(v1, w)
			x2, y2 := getCoordinates(v2, w)

			// If the edge is horizontal
			if y1 == y2 {
				if x2 < x1 {
					x1, y1, x2, y2 = x2, y2, x1, y1
				}
				maze[y1*3+1][x1*3+2] = true
				maze[y2*3+1][x2*3] = true
			} else {
				if y2 < y1 {
					x1, y1, x2, y2 = x2, y2, x1, y1
				}
				maze[y1*3+2][x1*3+1] = true
				maze[y2*3][x2*3+1] = true
			}
			callback()
		}
	}

	return maze

}

type intSet map[int]struct{}

func newIntSet() intSet {
	return make(intSet)
}

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

func getCoordinates(v, w int) (int, int) {
	return v % w, v / w
}

// getNeighbours returns the top, right, bottom, and left neighbors of a vertex
func getNeighbours(v, w int) (int, int, int, int) {
	top, right, bottom, left := -1, -1, -1, -1
	verticesInRow := w
	rowIdx := v / verticesInRow
	colIdx := v % verticesInRow
	if rowIdx > 0 {
		top = v - verticesInRow
	}
	if colIdx > 0 {
		left = v - 1
	}
	if rowIdx < w-1 {
		bottom = v + verticesInRow
	}
	if colIdx < w-1 {
		right = v + 1
	}
	return top, right, bottom, left
}

// getVertices returns the two vertices connected by the given edge
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

func (s intSet) add(i int) {
	s[i] = struct{}{}
}

func (s intSet) remove(i int) {
	delete(s, i)
}

func (s intSet) contains(i int) bool {
	_, ok := s[i]
	return ok
}

func newUnionSet(edges int) *unionSet {
	uf := &unionSet{
		parents: make([]int, edges),
		ranks:   make([]int, edges),
	}
	for i := range uf.parents {
		uf.parents[i] = i
		uf.ranks[i] = 0
	}
	return uf
}

func (u *unionSet) connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *unionSet) find(p int) int {
	if u.parents[p] == p {
		return p
	}
	return u.find(u.parents[p])
}

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

type unionSet struct {
	parents []int
	ranks   []int
}
