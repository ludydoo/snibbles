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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBellmanFord(t *testing.T) {

	a := 0
	b := 1
	c := 2
	d := 3
	e := 4
	f := 5

	dist, pred := BellmanFord([][]int{
		{a, c, 20},
		{a, b, 10},
		{b, e, 10},
		{b, d, 50},
		{c, d, 20},
		{c, e, 33},
		{d, e, -20},
		{e, f, 1},
		{d, f, -2},
	}, a)

	expectDist := []int{0, 10, 20, 40, 20, 21}
	expectPred := []int{-1, 0, 0, 2, 1, 4}
	assert.Equal(t, expectDist, dist)
	assert.Equal(t, expectPred, pred)

}
