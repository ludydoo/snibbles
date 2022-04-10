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
	"github.com/stretchr/testify/assert"
	"testing"
)

type option func(node *Node[int])

func rightNode(right *Node[int]) option {
	return func(node *Node[int]) {
		node.Right = right
		right.Parent = node
	}
}

func leftNode(left *Node[int]) option {
	return func(node *Node[int]) {
		node.Left = left
		left.Parent = node
	}
}

func aNode(key int, options ...option) *Node[int] {
	n := &Node[int]{Key: key}
	for _, option := range options {
		option(n)
	}
	return n
}

func TestNode_MinNode(t *testing.T) {

	tests := []struct {
		name string
		node *Node[int]
		want *Node[int]
	}{
		{
			name: "nil node",
			node: nil,
			want: nil,
		},
		{
			name: "no children",
			node: aNode(1),
			want: aNode(1),
		}, {
			name: "left node",
			node: aNode(2, leftNode(aNode(1))),
			want: aNode(1),
		}, {
			name: "right node",
			node: aNode(2, rightNode(aNode(3))),
			want: aNode(2),
		}, {
			name: "left child node",
			node: aNode(3,
				leftNode(
					aNode(2,
						leftNode(
							aNode(1),
						),
					),
				),
			),
			want: aNode(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wantKey *int
			if tt.want != nil {
				wantKey = &tt.want.Key
			}
			var actualKey *int
			actualNode := tt.node.MinNode()
			if actualNode != nil {
				actualKey = &actualNode.Key
			}
			assert.Equalf(t, wantKey, actualKey, "MinNode()")
		})
	}
}

func TestNode_MaxNode(t *testing.T) {

	tests := []struct {
		name string
		node *Node[int]
		want *Node[int]
	}{
		{
			name: "nil node",
			node: nil,
			want: nil,
		},
		{
			name: "no children",
			node: aNode(1),
			want: aNode(1),
		}, {
			name: "left node",
			node: aNode(2, leftNode(aNode(1))),
			want: aNode(2),
		}, {
			name: "right node",
			node: aNode(2, rightNode(aNode(3))),
			want: aNode(3),
		}, {
			name: "right child node",
			node: aNode(1,
				rightNode(
					aNode(2,
						rightNode(
							aNode(3),
						),
					),
				),
			),
			want: aNode(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wantKey *int
			if tt.want != nil {
				wantKey = &tt.want.Key
			}
			var actualKey *int
			actualNode := tt.node.MaxNode()
			if actualNode != nil {
				actualKey = &actualNode.Key
			}
			assert.Equalf(t, wantKey, actualKey, "MaxNode()")
		})
	}
}
