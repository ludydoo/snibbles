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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIslands(t *testing.T) {

	testCases := []struct {
		name   string
		mat    [][]int
		expect [][]int
	}{
		{
			name:   "zero",
			mat:    [][]int{},
			expect: [][]int{},
		}, {
			name: "1x1",
			mat: [][]int{
				{1},
			},
			expect: [][]int{
				{1},
			},
		}, {
			name: "2x2",
			mat: [][]int{
				{0, 0},
				{0, 1},
			},
			expect: [][]int{
				{0, 0},
				{0, 1},
			},
		}, {
			name: "simple",
			mat: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expect: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		}, {
			name: "full",
			mat: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			expect: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
		}, {
			name: "line",
			mat: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
			},
			expect: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
			},
		}, {
			name: "funk",
			mat: [][]int{
				{1, 0, 0, 0, 1},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 1, 1},
				{1, 1, 0, 0, 1},
				{1, 0, 1, 0, 1},
			},
			expect: [][]int{
				{1, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1},
				{1, 1, 0, 0, 1},
				{1, 0, 1, 0, 1},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := RemoveIslands(testCase.mat)
			assert.Equal(t, testCase.expect, actual)
		})
	}

}
