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

package prim

import "math"

func PrimsAlgorithm(mat [][]int) [][]int {

	var (
		// number of vertices
		numNodes int

		// store the cost of the edges
		cost [][]int

		// store visited nodes
		visited []bool

		// stores the edge with the lowest cost
		// the index of the array is the index of the predecessor
		// the value of the array is the index of the successor
		predecessors []int
	)

	numNodes = len(mat)
	visited = make([]bool, numNodes)
	predecessors = make([]int, numNodes)
	cost = make([][]int, numNodes)
	visited[0] = true

	// init cost matrix
	for i := 0; i < numNodes; i++ {
		cost[i] = make([]int, numNodes)
		for j := 0; j < numNodes; j++ {
			cost[i][j] = mat[i][j]
			// if i == j, set cost to infinity
			if mat[i][j] == 0 {
				cost[i][j] = math.MaxInt32
			}
		}
	}

	// loop on all nodes > 0
	// because we start from node 0, and we want to visit only the other nodes
	for k := 1; k < numNodes; k++ {

		// x is the node with the minimum cost
		// y is the node with the minimum cost from x
		var x, y int

		// find the minimum cost edge from any visited node to any unvisited node
		// store the source of the edge in x and the destination in y

		// loop on all nodes
		for i := 0; i < numNodes; i++ {
			for j := 0; j < numNodes; j++ {
				// if the node was not visited, and the cost is lower than the current one
				// update the current node
				if visited[i] && !visited[j] && cost[i][j] < cost[x][y] {
					// x is the visited node
					x = i
					// y is the unvisited node
					y = j
				}
			}
		}
		println("next selected edge: ", x, ",", y, " cost: ", cost[x][y])
		// set the parent of y to x
		predecessors[y] = x
		// mark y as visited
		visited[y] = true
	}

	// print the edges
	for i := 0; i < numNodes; i++ {
		println(predecessors[i], "-->", i)
	}

	// build the result edges
	var result [][]int
	for i := 0; i < numNodes; i++ {
		result = append(result, []int{predecessors[i], i})
	}

	return result

}
