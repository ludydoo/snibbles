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
	"dsa/pkg/stack"
	"golang.org/x/exp/constraints"
)

func Traverse[V constraints.Ordered](n *Node[V], mode Traversal, fn func(v V) bool) error {
	switch mode {
	case PreOrder:
		PreOrderTraversal(n, fn)
	case InOrder:
		InOrderTraversal(n, fn)
	case PostOrder:
		PostOrderTraversal(n, fn)
	case LevelOrder:
		LevelOrderTraversal(n, fn)
	default:
		return ErrInvalidTraversalType{
			msg: "invalid traversal type",
		}
	}
	return nil
}

func Elements[v constraints.Ordered](n *Node[v], mode Traversal) ([]v, error) {
	var elements []v
	err := Traverse(n, mode, func(v v) bool {
		elements = append(elements, v)
		return true
	})
	return elements, err
}

// InOrderTraversal traverses the tree in-order
func InOrderTraversal[V constraints.Ordered](n *Node[V], fn func(v V) bool) {
	s := stack.New[*Node[V]]()
	cur := n
	for cur != nil || !s.IsEmpty() {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			cur, _ = s.Pop()
			if !fn(cur.Key) {
				break
			}
			cur = cur.Right
		}
	}
}

// PreOrderTraversal traverses the tree pre-order
func PreOrderTraversal[V constraints.Ordered](n *Node[V], fn func(v V) bool) {
	if n == nil {
		return
	}
	s := stack.New[*Node[V]]()
	s.Push(n)
	for !s.IsEmpty() {
		cur, _ := s.Pop()
		if !fn(cur.Key) {
			break
		}
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

// PostOrderTraversal traverses the tree post-order iteratively
func PostOrderTraversal[V constraints.Ordered](n *Node[V], fn func(v V) bool) {
	if n == nil {
		return
	}
	s := stack.New[*Node[V]]()
	var last *Node[V] = nil
	cur := n
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			peek := s.Peek()
			if peek.Right != nil && last != peek.Right {
				cur = peek.Right
			} else {
				if !fn(peek.Key) {
					break
				}
				last, _ = s.Pop()
			}
		}
	}
}

func LevelOrderTraversal[V constraints.Ordered](n *Node[V], fn func(v V) bool) {
	if n == nil {
		return
	}
	queue := []*Node[V]{n}
	for len(queue) > 0 {
		node := queue[0]
		if !fn(node.Key) {
			break
		}
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
