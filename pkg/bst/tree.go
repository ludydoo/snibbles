package bst

import (
	"golang.org/x/exp/constraints"
	"strings"
)

// Tree is a Binary Search Tree
type Tree[V constraints.Ordered] struct {
	// Root Node
	Root *Node[V]
}

// New creates a new tree with the given items
func New[V constraints.Ordered](items ...V) *Tree[V] {
	tree := &Tree[V]{}
	for _, item := range items {
		tree.Insert(item)
	}
	return tree
}

func (t *Tree[V]) Traversal(mode Traversal, fn func(v V) bool) error {
	return Traverse(t.Root, mode, fn)
}

// InOrderTraversal traverses the tree in-order
func (t *Tree[V]) InOrderTraversal(fn func(v V) bool) error {
	return t.Traversal(InOrder, fn)
}

// InOrderIterator returns an in-order iterator
func (t *Tree[V]) InOrderIterator() *InOrderIterator[V] {
	iterator := newInOrderIterator[V](t.Root)
	return iterator
}

// InOrder returns elements in-order
func (t *Tree[V]) InOrder() ([]V, error) {
	return Elements(t.Root, InOrder)
}

// PreOrderTraversal traverses the tree pre-order
func (t *Tree[V]) PreOrderTraversal(fn func(v V) bool) error {
	return t.Traversal(PreOrder, fn)
}

// PreOrder returns elements pre-order
func (t *Tree[V]) PreOrder() ([]V, error) {
	return Elements(t.Root, PreOrder)
}

// PostOrderTraversal traverses the tree post-order
func (t *Tree[V]) PostOrderTraversal(fn func(v V) bool) error {
	return t.Traversal(PostOrder, fn)
}

// PostOrder returns elements post-order
func (t *Tree[V]) PostOrder() ([]V, error) {
	return Elements(t.Root, PostOrder)
}

// LevelOrderTraversal traverses the tree level-order
func (t *Tree[V]) LevelOrderTraversal(fn func(v V) bool) error {
	return t.Traversal(LevelOrder, fn)
}

// LevelOrder returns elements level-ordered
func (t *Tree[V]) LevelOrder() ([]V, error) {
	return Elements(t.Root, LevelOrder)
}

// Delete an item
func (t *Tree[V]) Delete(v V) {
	if t.Root == nil {
		return
	}
	t.Root = t.Root.Remove(v)
}

// Insert an item
func (t *Tree[V]) Insert(v V) {
	if t.Root == nil {
		t.Root = &Node[V]{Key: v}
	} else {
		t.Root.Insert(v)
	}
}

// Contains an item
func (t *Tree[V]) Contains(v V) bool {
	if t.Root == nil {
		return false
	}
	return t.Root.Contains(v)
}

// Render the tree
func (t *Tree[V]) Render() string {
	str := &strings.Builder{}
	str.WriteString("\n")
	str.WriteString(t.Root.str(1))
	str.WriteString("\n")
	return str.String()
}

func balance[V constraints.Ordered](keys []V, start, end int) *Node[V] {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := &Node[V]{Key: keys[mid]}
	node.Left = balance(keys, start, mid-1)
	if node.Left != nil {
		node.Left.Parent = node
	}
	node.Right = balance(keys, mid+1, end)
	if node.Right != nil {
		node.Right.Parent = node
	}
	return node
}

func (t *Tree[V]) Balance() error {
	inOrder, _ := t.InOrder()
	t.Root = balance(inOrder, 0, len(inOrder)-1)
	return nil
}
