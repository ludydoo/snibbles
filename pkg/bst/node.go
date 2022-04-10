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
	"golang.org/x/exp/constraints"
	"strings"
)

// Node represents a Tree Node
type Node[V constraints.Ordered] struct {
	// Key attached to the node
	Key V
	// Left Node
	Left *Node[V]
	// Right Node
	Right *Node[V]
	// Parent Node
	Parent *Node[V]
}

// MinNode find the node with the minimum key
func (n *Node[V]) MinNode() *Node[V] {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}
	return n.Left.MinNode()
}

// MaxNode find the node with the maximum key
func (n *Node[V]) MaxNode() *Node[V] {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n
	}
	return n.Right.MaxNode()
}

// remove is an internal method to remove a node with the specified key
// returns the new node
func (n *Node[V]) remove(v V) *Node[V] {
	if n == nil {
		return nil
	}
	if v < n.Key {
		n.Left = n.Left.remove(v)
		if n.Left != nil {
			n.Left.Parent = n
		}
		return n
	}
	if v > n.Key {
		n.Right = n.Right.remove(v)
		if n.Right != nil {
			n.Right.Parent = n
		}
		return n
	}
	if n.Left == nil && n.Right == nil {
		return nil
	}
	if n.Left == nil {
		return n.Right
	}
	if n.Right == nil {
		return n.Left
	}
	minNode := n.Right.MinNode()
	n.Key = minNode.Key
	n.Right = n.Right.remove(minNode.Key)
	if n.Right != nil {
		n.Right.Parent = n
	}
	return n
}

// Remove a key
func (n *Node[V]) Remove(v V) *Node[V] {
	return n.remove(v)
}

// str renders a node
func (n *Node[V]) str(level int) string {
	if n == nil {
		return ""
	}
	writer := &strings.Builder{}

	writer.WriteString(n.Right.str(level + 1))

	for i := 0; i < level; i++ {
		writer.WriteString("       ")
	}

	if n.Parent != nil && n.Parent.Right == n {
		writer.WriteString("┌── ")
	} else if n.Parent != nil && n.Parent.Left == n {
		writer.WriteString("└── ")
	} else {
		writer.WriteString("────")
	}

	writer.WriteString(fmt.Sprintf("%v ", n.Key))

	if n.Left != nil && n.Right != nil {
		writer.WriteString("┤")
	} else if n.Left != nil {
		writer.WriteString("┐")
	} else if n.Right != nil {
		writer.WriteString("┘")
	} else {
		writer.WriteString("")
	}
	writer.WriteString("\n")

	writer.WriteString(n.Left.str(level + 1))
	return writer.String()
}

// Insert a key
func (n *Node[V]) Insert(v V) {
	if n == nil {
		return
	}
	if v < n.Key {
		if n.Left == nil {
			n.Left = &Node[V]{Key: v, Parent: n}
		} else {
			n.Left.Insert(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node[V]{Key: v, Parent: n}
		} else {
			n.Right.Insert(v)
		}
	}
}

// Contains returns if the key exists
func (n *Node[V]) Contains(v V) bool {
	if n == nil {
		return false
	}
	if v == n.Key {
		return true
	}
	if v < n.Key {
		return n.Left.Contains(v)
	} else {
		return n.Right.Contains(v)
	}
}
