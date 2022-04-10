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

package bellmanford

import (
	"fmt"
	"math"
)

type intSet map[int]struct{}

func (s intSet) add(i int) {
	s[i] = struct{}{}
}

func (s intSet) len() int {
	return len(s)
}

func newIntSet() intSet {
	return make(map[int]struct{})
}

func BellmanFord(edges [][]int, source int) ([]int, []int) {

	uniqueVs := newIntSet()

	for _, edge := range edges {
		v1 := edge[0]
		v2 := edge[1]
		uniqueVs.add(v1)
		uniqueVs.add(v2)
	}

	vertexCount := uniqueVs.len()

	distance := make([]int, vertexCount)
	predecessor := make([]int, vertexCount)

	for i := 0; i < vertexCount; i++ {
		distance[i] = math.MaxInt32
		predecessor[i] = -1
	}

	distance[source] = 0

	for k := 1; k < vertexCount; k++ {
		for _, edge := range edges {
			v1 := edge[0]
			v2 := edge[1]
			w := edge[2]
			if distance[v1]+w < distance[v2] {
				distance[v2] = distance[v1] + w
				predecessor[v2] = v1
			}
		}
	}
	for _, edge := range edges {
		v1 := edge[0]
		v2 := edge[1]
		w := edge[2]
		if distance[v1]+w < distance[v2] {
			panic("graph contains a negative-weight cycle")
		}
	}

	for i := range predecessor {
		fmt.Printf("%d -> %d\n", predecessor[i], i)
	}

	for i, dist := range distance {
		fmt.Printf("%d - %d\n", i, dist)
	}

	return distance, predecessor
}
