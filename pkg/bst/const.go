package bst

type Traversal uint

const (
	PreOrder Traversal = iota
	InOrder
	PostOrder
	LevelOrder
)

type Direction uint

const (
	Forward Direction = iota
	Backward
)
