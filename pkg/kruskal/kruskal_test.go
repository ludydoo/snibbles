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

package kruskal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKruskal(t *testing.T) {

	graph := []Edge{
		{1, 2, 2},
		{1, 4, 1},
		{1, 5, 4},
		{5, 4, 9},
		{4, 3, 5},
		{4, 2, 3},
		{2, 3, 3},
		{2, 6, 7},
		{6, 3, 8},
	}

	expected := []Edge{
		{1, 4, 1},
		{1, 2, 2},
		{2, 3, 3},
		{1, 5, 4},
		{2, 6, 7},
	}

	print("Input: \n")
	printGraph(graph)

	result := Kruskal(graph)

	print("Output: \n")
	printGraph(result)

	assert.True(t, graphsEqual(result, expected))
}
