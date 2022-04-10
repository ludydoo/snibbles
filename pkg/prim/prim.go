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
		// cost matrix
		cost [][]int
		// visited nodes
		visited []bool
		// min cost edges
		predNode []int
	)

	numNodes = len(mat)
	visited = make([]bool, numNodes)
	predNode = make([]int, numNodes)
	cost = make([][]int, numNodes)
	visited[0] = true

	for i := 0; i < numNodes; i++ {
		cost[i] = make([]int, numNodes)
		for j := 0; j < numNodes; j++ {
			cost[i][j] = mat[i][j]
			if mat[i][j] == 0 {
				cost[i][j] = math.MaxInt32
			}
		}
	}

	for k := 1; k < numNodes; k++ {
		var x, y int
		for i := 0; i < numNodes; i++ {
			for j := 0; j < numNodes; j++ {
				if visited[i] && !visited[j] && cost[i][j] < cost[x][y] {
					x = i
					y = j
				}
			}
		}
		println("next selected edge: ", x, ",", y, " cost: ", cost[x][y])
		predNode[y] = x
		visited[y] = true
	}

	for i := 0; i < numNodes; i++ {
		println(predNode[i], "-->", i)
	}

	var result [][]int
	for i := 0; i < numNodes; i++ {
		result = append(result, []int{predNode[i], i})
	}

	return result

}
