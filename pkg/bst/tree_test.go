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

package bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"sort"
	"strconv"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	tree := &Tree[int]{}
	for i := 0; i < 10; i++ {
		tree.Insert(rand.Intn(100))
	}
	assert.True(t, tree.Contains(10))
	assert.False(t, tree.Contains(200))
	print(tree.Render())
}

func TestTree_Contains(t *testing.T) {
	tc := []struct {
		name     string
		tree     *Tree[int]
		value    int
		expected bool
	}{
		{
			name:     "empty tree",
			tree:     &Tree[int]{},
			value:    10,
			expected: false,
		}, {
			name:     "tree with one element",
			tree:     New(10),
			value:    10,
			expected: true,
		}, {
			name:     "tree with one element - does not contain",
			tree:     New(10),
			value:    20,
			expected: false,
		}, {
			name:     "tree with multiple elements",
			tree:     New(10, 20, 30, 40),
			value:    40,
			expected: true,
		}, {
			name:     "tree with multiple elements - does not contain",
			tree:     New(10, 20, 30, 40),
			value:    50,
			expected: false,
		},
	}
	for _, s := range tc {
		t.Run(s.name, func(t *testing.T) {
			assert.Equal(t, s.expected, s.tree.Contains(s.value))
		})
	}
}

func TestTree_InOrderIterator(t *testing.T) {

	var intsInOrder []int
	var intsReverseOrder []int
	for i := 0; i < 10; i++ {
		rantInt := rand.Intn(100)
		intsInOrder = append(intsInOrder, rantInt)
	}

	tree := New[int](intsInOrder...)
	sort.Ints(intsInOrder)
	intsReverseOrder = reverseArray(intsInOrder)

	t.Log(tree.Render())

	it := tree.InOrderIterator()
	var inOrder []int
	var reverseOrder []int

	for it.Next() {
		t.Log(it.Value())
		inOrder = append(inOrder, it.Value())
	}

	for it.Previous() {
		t.Log(it.Value())
		reverseOrder = append(reverseOrder, it.Value())
	}

	assert.Equal(t, intsInOrder, inOrder)
	assert.Equal(t, intsReverseOrder, reverseOrder)
}

func TestTree_InOrder(t *testing.T) {
	expected := []int{30, 40, 45, 50, 55, 60, 70}
	tree := New(50, 40, 60, 70, 30, 55, 45)
	actual, err := tree.InOrder()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, expected, actual)
}

func TestTree_PostOrder(t *testing.T) {
	tree := New(50, 40, 60, 70, 30, 55, 45)
	expected := []int{30, 45, 40, 55, 70, 60, 50}
	actual, err := tree.PostOrder()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, expected, actual)
}

func TestTree_PreOrder(t *testing.T) {
	expected := []int{50, 40, 30, 45, 60, 55, 70}
	tree := New(50, 40, 60, 70, 30, 55, 45)
	actual, err := tree.PreOrder()
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, expected, actual)
}

func TestTree_LevelOrderTraversal(t *testing.T) {
	tree := New(50, 40, 60, 70, 30, 55, 45)
	fmt.Println(tree.Render())
	var items []int
	err := tree.LevelOrderTraversal(func(value int) bool {
		fmt.Printf("%v\n", value)
		items = append(items, value)
		return true
	})
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, []int{50, 40, 60, 30, 45, 55, 70}, items)
}

func TestTree_Balance(t *testing.T) {
	tree := New(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(tree.Render())
	err := tree.Balance()
	if !assert.NoError(t, err) {
		return
	}
	fmt.Println(tree.Render())
	assert.Equal(t, New(4, 2, 6, 1, 3, 5, 7), tree)
}

func TestTree_Delete_Fuzz(t *testing.T) {
	randInts := map[string]int{}
	for i := 0; i < 400; i++ {
		randInts[strconv.Itoa(i)] = i
	}
	var ints []int
	for _, i := range randInts {
		ints = append(ints, i)
	}
	tree := New(ints...)
	for _, i := range ints {
		tree.Delete(i)
		ints = remove(ints, i)
		inOrder, err := tree.InOrder()
		if !assert.NoError(t, err) {
			return
		}
		assert.ElementsMatch(t, ints, inOrder)
	}
}

func Benchmark_Tree_Delete(b *testing.B) {
	randInts := map[string]int{}
	for len(randInts) < b.N {
		randInt := rand.Intn(b.N * 10)
		randInts[strconv.Itoa(rand.Intn(b.N*10))] = randInt
	}
	var ints []int
	for _, i := range randInts {
		ints = append(ints, i)
	}
	b.ResetTimer()
	tree := New(ints...)
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}

func remove(arr []int, without int) []int {
	var newArr []int
	for _, i := range arr {
		if i != without {
			newArr = append(newArr, i)
		}
	}
	return newArr
}

func TestTree_Delete(t *testing.T) {

	tcs := []struct {
		name string
		tree *Tree[int]
		key  int
		exp  *Tree[int]
	}{
		{
			name: "delete with 2 children",
			tree: New(50, 40, 60, 30, 45, 55, 70),
			key:  40,
			exp:  New(50, 45, 60, 30, 55, 70),
		},
		{
			name: "delete root",
			tree: New(50, 40, 60, 30, 45, 55, 70),
			key:  50,
			exp:  New(55, 40, 60, 30, 45, 70),
		},
		{
			name: "delete leaf",
			tree: New(50, 40, 60, 30, 45, 55, 70),
			key:  45,
			exp:  New(50, 40, 60, 30, 55, 70),
		},
		{
			name: "delete with 1 child",
			tree: New(50, 40, 60, 30, 45, 55),
			key:  60,
			exp:  New(50, 40, 55, 30, 45),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("\n%v\n", tc.tree.Render())
			tc.tree.Delete(tc.key)
			t.Logf("\n%v\n", tc.tree.Render())
			assert.Equal(t, tc.exp, tc.tree)
		})
	}
}

func reverseArray(arr []int) []int {
	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
