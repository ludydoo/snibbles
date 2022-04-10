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
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetVertices(t *testing.T) {

	testCases := []struct {
		name   string
		w      int
		expect [][]int
	}{
		{
			// 0 - 1
			// |   |
			// 2 - 3
			name: "2x2",
			w:    2,
			expect: [][]int{
				{0, 1},
				{0, 2},
				{1, 3},
				{2, 3},
			},
		}, {
			// 0 - 1 - 2
			// |   |   |
			// 3 - 4 - 5
			// |   |   |
			// 6 - 7 - 8
			name: "3x3",
			w:    3,
			expect: [][]int{
				{0, 1},
				{1, 2},
				{0, 3},
				{1, 4},
				{2, 5},
				{3, 4},
				{4, 5},
				{3, 6},
				{4, 7},
				{5, 8},
				{6, 7},
				{7, 8},
			},
		}, {
			// 0 - 1 - 2 - 3
			// |   |   |   |
			// 4 - 5 - 6 - 7
			// |   |   |   |
			// 8 - 9 - 10- 11
			// |   |   |   |
			// 12- 13- 14- 15
			// |   |   |   |
			// 16- 17- 18- 19
			name: "4x5",
			w:    4,
			expect: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{0, 4},
				{1, 5},
				{2, 6},
				{3, 7},
				{4, 5},
				{5, 6},
				{6, 7},
				{4, 8},
				{5, 9},
				{6, 10},
				{7, 11},
				{8, 9},
				{9, 10},
				{10, 11},
				{8, 12},
				{9, 13},
				{10, 14},
				{11, 15},
				{12, 13},
				{13, 14},
				{14, 15},
				{12, 16},
				{13, 17},
				{14, 18},
				{15, 19},
				{16, 17},
				{17, 18},
				{18, 19},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for i, ints := range testCase.expect {
				expect1 := ints[0]
				expect2 := ints[1]
				t.Run(testCase.name+"_"+strconv.Itoa(i), func(t *testing.T) {
					v1, v2 := getVertices(i, testCase.w)

					fmt.Printf("Actual: %d %d\n", v1, v2)
					fmt.Printf("Expect: %d %d\n", expect1, expect2)

					assert.Equalf(t, expect1, v1, "First Edge")
					assert.Equalf(t, expect2, v2, "Second Edge")
				})
			}
		})
	}

}

func TestGetNeighbours(t *testing.T) {
	testCases := []struct {
		name   string
		w      int
		expect [][]int
	}{
		{
			// 0 - 1 - 2 - 3
			// |   |   |   |
			// 4 - 5 - 6 - 7
			// |   |   |   |
			// 8 - 9 - 10 - 11
			// |   |   |   |
			// 12- 13- 14- 15
			name: "4x4",
			w:    4,
			expect: [][]int{
				{-1, 1, 4, -1},
				{-1, 2, 5, 0},
				{-1, 3, 6, 1},
				{-1, -1, 7, 2},
				{0, 5, 8, -1},
				{1, 6, 9, 4},
				{2, 7, 10, 5},
				{3, -1, 11, 6},
				{4, 9, 12, -1},
				{5, 10, 13, 8},
				{6, 11, 14, 9},
				{7, -1, 15, 10},
				{8, 13, -1, -1},
				{9, 14, -1, 12},
				{10, 15, -1, 13},
				{11, -1, -1, 14},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for expectIdx, expect := range testCase.expect {
				t.Run(strconv.Itoa(expectIdx), func(t *testing.T) {

					top, right, bottom, left := getNeighbours(expectIdx, testCase.w)
					expectTop := expect[0]
					expectRight := expect[1]
					expectBottom := expect[2]
					expectLeft := expect[3]

					t.Logf("Case #%d", expectIdx)
					t.Logf("Top. Expect %d. Got %d", expectTop, top)
					t.Logf("Right. Expect %d. Got %d", expectRight, right)
					t.Logf("Bottom. Expect %d. Got %d", expectBottom, bottom)
					t.Logf("Left. Expect %d. Got %d", expectLeft, left)

					assert.Equal(t, expectTop, top, "Top")
					assert.Equal(t, expectRight, right, "Right")
					assert.Equal(t, expectBottom, bottom, "Bottom")
					assert.Equal(t, expectLeft, left, "Left")
				})
			}
		})
	}
}

func TestNewMaze(t *testing.T) {
	NewMaze(40, 20)
}
