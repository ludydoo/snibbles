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
