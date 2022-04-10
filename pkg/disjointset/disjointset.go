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

package disjointset

// Set represents a disjoint set.
type Set[V comparable] struct {
	// parents is a map of the parent of each node
	parents map[V]V
	// ranks is a map of the rank of each node
	ranks map[V]int
}

// NewSet creates a new disjoint set.
func NewSet[V comparable]() *Set[V] {
	return &Set[V]{
		parents: make(map[V]V),
		ranks:   make(map[V]int),
	}
}

// MakeSet creates a new set with the given node as the only member.
func MakeSet[V comparable](v V) *Set[V] {
	s := NewSet[V]()
	s.parents[v] = v
	s.ranks[v] = 0
	return s
}

// Find returns the representative of the set containing the given node.
func (s *Set[V]) Find(item V) V {
	if s.parents[item] == item {
		return item
	}
	return s.Find(s.parents[item])
}

// Union merges the sets containing the given items.
func (s *Set[V]) Union(item1, item2 V) {
	root1 := s.Find(item1)
	root2 := s.Find(item2)
	if root1 == root2 {
		return
	}
	if s.ranks[root1] < s.ranks[root2] {
		s.parents[root1] = root2
	} else if s.ranks[root1] > s.ranks[root2] {
		s.parents[root2] = root1
	} else {
		s.parents[root2] = root1
		s.ranks[root1]++
	}
}
