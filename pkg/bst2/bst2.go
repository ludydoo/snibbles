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

package bst2

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

// Node is a binary tree Node holding a Key and a value.
type Node struct {
	// Key is the Node's Key.
	Key int
	// Left is the Node's Left child.
	Left *Node
	// Right is the Node's Right child.
	Right *Node
}

// Insert a new Node into the BST.
func Insert(root *Node, key int) *Node {
	cur := root
	if root == nil {
		return &Node{Key: key}
	}
	for {
		if key < cur.Key {
			if cur.Left == nil {
				cur.Left = &Node{Key: key}
				return root
			} else {
				cur = cur.Left
			}
		} else if key > cur.Key {
			if cur.Right == nil {
				cur.Right = &Node{Key: key}
				return root
			} else {
				cur = cur.Right
			}
		} else {
			return root
		}
	}
}

// TraverseInOrder traverses the BST in order.
func TraverseInOrder(root *Node) []int {
	var result []int
	cur := root
	stack := make([]*Node, 0)
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Key)
			cur = cur.Right
		}
	}
	return result
}

// TraversePreOrder traverses the BST in pre-order.
func TraversePreOrder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	cur := root
	stack := make([]*Node, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			result = append(result, cur.Key)
		}
		if cur != nil && cur.Right != nil {
			stack = append(stack, cur.Right)
			cur = cur.Left
		} else {
			if len(stack) > 0 {
				cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				cur = nil
			}
		}
	}
	return result
}

// TraversePostOrder traverses the BST in post-order.
func TraversePostOrder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	cur := root
	var last *Node
	stack := make([]*Node, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else if len(stack) > 0 {
			peek := stack[len(stack)-1]
			if peek.Right != nil && peek.Right != last {
				cur = peek.Right
			} else {
				last = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result = append(result, peek.Key)
				fmt.Println(peek.Key)
			}
		}
	}
	return result
}

// TraverseLevelOrder traverses the BST in level-order.
func TraverseLevelOrder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	q := []*Node{root}
	for len(q) > 0 {
		next := q[0]
		q = q[1:]
		result = append(result, next.Key)
		if next.Left != nil {
			q = append(q, next.Left)
		}
		if next.Right != nil {
			q = append(q, next.Right)
		}
	}
	return result
}

// TraverseReverseLevelOrder traverses the BST in reverse level-order.
func TraverseReverseLevelOrder(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	var queue []*Node
	var stack []*Node
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		stack = append(stack, node)
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Key)
	}
	return result
}

func GetParent(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if key < cur.Key {
			if cur.Left != nil && cur.Left.Key == key {
				return cur
			}
			cur = cur.Left
		} else if key > cur.Key {
			if cur.Right != nil && cur.Right.Key == key {
				return cur
			}
			cur = cur.Right
		} else {
			return nil
		}
	}
	return nil
}

func FindInOrderSuccessor(root *Node, key int) *Node {
	cur := root
	stack := []*Node{}

	for {
		if cur != nil {
			if cur.Key < key {
				stack = append(stack, cur)
				cur = cur.Right
			} else if cur.Key > key {
				stack = append(stack, cur)
				cur = cur.Left
			} else {
				if cur.Right != nil {
					return MinNode(cur.Right)
				} else {
					for {
						if len(stack) == 0 {
							return nil
						}
						parent := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						if parent.Left == cur {
							return parent
						}
						cur = parent
					}
				}
			}
		}
	}
}

func FindInOrderPredecessor(root *Node, key int) *Node {
	cur := root
	stack := []*Node{}
	for {
		if cur != nil {
			if cur.Key < key {
				stack = append(stack, cur)
				cur = cur.Right
			} else if cur.Key > key {
				stack = append(stack, cur)
				cur = cur.Left
			} else {
				if cur.Left != nil {
					return MaxNode(cur.Left)
				} else {
					for {
						if len(stack) == 0 {
							return nil
						}
						parent := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						if parent.Right == cur {
							return parent
						}
						cur = parent
					}
				}
			}
		}
	}
}

// PrintFromLeft prints the BST viewed from the Left.
func PrintFromLeft(root *Node) []int {
	var result []int
	if root == nil {
		return result
	}
	q := []*Node{root}
	for len(q) > 0 {
		queueLen := len(q)
		for i := 0; i < queueLen; i++ {
			queueItem := q[i]
			if i == 0 {
				result = append(result, queueItem.Key)
			}
			if queueItem.Left != nil {
				q = append(q, queueItem.Left)
			}
			if queueItem.Right != nil {
				q = append(q, queueItem.Right)
			}
		}
		q = q[queueLen:]
	}
	return result
}

func DeleteKey(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key < root.Key {
		root.Left = DeleteKey(root.Left, key)
	} else if key > root.Key {
		root.Right = DeleteKey(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		}
		successor := MinNode(root.Right)
		root.Key = successor.Key
		root.Right = DeleteKey(root.Right, root.Key)
	}
	return root
}

// MinNode returns the Node with the minimum value
func MinNode(node *Node) *Node {
	cur := node
	for {
		if cur.Left != nil {
			cur = cur.Left
		} else {
			return cur
		}
	}
}

// MaxNode returns the Node with the maximum value
func MaxNode(node *Node) *Node {
	cur := node
	for {
		if cur.Right != nil {
			cur = cur.Right
		} else {
			return cur
		}
	}
}

// IsSymmetric checks if the BST is symmetric
func IsSymmetric(root *Node) bool {
	queue := []*Node{root}
	for {
		queueLen := len(queue)
		if queueLen == 0 {
			return true
		}
		for i := 0; i < queueLen/2; i++ {
			if (queue[i] == nil) != (queue[queueLen-i-1] == nil) {
				return false
			}
		}
		for _, item := range queue {
			if item != nil {
				queue = append(queue, item.Left)
				queue = append(queue, item.Right)
			}
		}
		queue = queue[queueLen:]
	}
}

// New returns a new BST
func New(items ...int) *Node {
	if len(items) == 0 {
		return nil
	}
	root := &Node{Key: items[0]}
	for i := 1; i < len(items); i++ {
		root = Insert(root, items[i])
	}
	return root
}

type printCallback func(node *Node) string
type printOptions struct {
	callback printCallback
}

type PrintOption func(o *printOptions)

func PrintCallback(callback printCallback) PrintOption {
	return func(o *printOptions) {
		o.callback = callback
	}
}

// Print prints the BST
// Will output something like
//
//             50
//    ╭────────┴────────╮
//    30                70
// ╭──┴──╮        ╭─────┴────────╮
// 20    40       60             80
//       └──╮     └──╮     ╭─────┴──╮
//          45       65    78       85
//                         └──╮     └──╮
//                            79       100
//                                     └───╮
//                                         200
//
func Print(root *Node, options ...PrintOption) string {
	opts := &printOptions{}
	for _, option := range options {
		option(opts)
	}
	if root == nil {
		return ""
	}
	result, _ := printNodeInt(root, opts)
	return result
}

// printNodeInt internal method for printing the BST
func printNodeInt(root *Node, opts *printOptions) (string, int) {

	//
	// lW : width of the Left part
	// rW : width of the Right part
	// lPad : Left hook index
	// rPad : Right hook index
	// R: root Node
	// lA: Left arm
	// rA: Right arm
	// rM: merge position for the Right part
	//
	// lA = lLen - lPAd
	// rA = rLen - rPad
	// rM = lLen + 3
	//
	//
	//       ╭ lA ┤R├ rA ──╮
	//       │             │
	// ├───lLen───┤ ├─────rLen─────┤
	//    ╭──╯             ╰───╮
	// lPad                  rPad
	//
	//              rM

	left := root.Left
	right := root.Right
	keyIntf := root.Key

	var rootVal string

	if opts.callback != nil {
		rootVal = opts.callback(root)
	}

	if rootVal == "" {
		rootVal = fmt.Sprintf("%d", keyIntf)
	}

	hasLeft := reflect.ValueOf(left).IsNil() == false
	hasRight := reflect.ValueOf(right).IsNil() == false

	if !hasLeft && !hasRight {
		return rootVal, 0
	}

	var leftVal, rightVal string
	var leftPad, rightPad int
	var leftArm, rightArm int

	if hasLeft {
		leftVal, leftPad = printNodeInt(left, opts)
	}
	if hasRight {
		rightVal, rightPad = printNodeInt(right, opts)
	}

	leftLen := 0
	leftRows := strings.Split(leftVal, "\n")
	for _, row := range leftRows {
		charCount := utf8.RuneCountInString(row)
		if charCount > leftLen {
			leftLen = charCount
		}
	}

	if hasLeft {
		leftArm = maxInt(1, leftLen-leftPad)
	}
	if hasRight {
		rightArm = maxInt(len(rootVal), rightPad+len(rootVal))
	}

	b := &strings.Builder{}
	if hasLeft {
		b.WriteString(strings.Repeat(" ", leftLen+1))
	}

	b.WriteString(rootVal)
	if hasLeft || hasRight {
		b.WriteString("\n")
	}

	if hasLeft {
		b.WriteString(strings.Repeat(" ", leftPad))
		b.WriteString("╭")
		b.WriteString(strings.Repeat("─", leftArm))
	} else {
		b.WriteString(strings.Repeat(" ", leftArm))
	}

	if hasLeft && hasRight {
		b.WriteString("┴")
	} else if hasLeft {
		b.WriteString("┘")
	} else if hasRight {
		b.WriteString("└")
	}

	if hasRight {
		b.WriteString(strings.Repeat("─", rightArm))
		b.WriteString("╮")
	} else {
		b.WriteString(strings.Repeat(" ", rightArm))
	}

	str := b.String()
	var merged string
	var resultInt int
	if !hasLeft && hasRight {
		merged = mergeStrs(leftVal, rightVal, 1+len(rootVal), 0)
		resultInt = 0
	} else if hasLeft && !hasRight {
		merged = mergeStrs(leftVal, rightVal, len(rootVal), 0)
		resultInt = leftLen + 1
	} else {
		merged = mergeStrs(leftVal, rightVal, leftLen+2+len(rootVal), 0)
		resultInt = maxInt(leftLen+1, 3)
	}
	merged = mergeStrs(str, merged, 0, 2)
	return merged, resultInt

}

// mergeStrs merges two strings based on the given offsets
// used to print BSTs
// will merge string like so
//
//    ╭────╮     ╭─────╮                   ╭──────╮
//    │1111│     │22222│    (x=1, y=2)     │1111  │
//    │1111│  x  │22222│                   │1111  │
//    │1111│     │22222│                   │122222│
//    ╰────╯     │22222│                   │ 22222│
//               ╰─────╯                   │ 22222│
//                                         │ 22222│
//                                         ╰──────╯
func mergeStrs(s1, s2 string, x, y int) string {

	s1Rows := strings.Split(s1, "\n")
	s2Rows := strings.Split(s2, "\n")

	res := strings.Builder{}

	rows := maxInt(len(s1Rows), len(s2Rows)+absInt(y))
	for r := 0; r < rows; r++ {

		var s1RowIdx int
		var s2RowIdx int
		var s1Index int
		var s2Index int
		if y < 0 {
			s1RowIdx = r + y
			s2RowIdx = r
		} else {
			s1RowIdx = r
			s2RowIdx = r - y
		}
		if x < 0 {
			s1Index = x
			s2Index = 0
		} else {
			s1Index = 0
			s2Index = -x
		}

		var s1Row []rune
		var s2Row []rune
		if s1RowIdx >= 0 && s1RowIdx < len(s1Rows) {
			s1Row = []rune(s1Rows[s1RowIdx])
		}
		if s2RowIdx >= 0 && s2RowIdx < len(s2Rows) {
			s2Row = []rune(s2Rows[s2RowIdx])
		}
		s1CharCount := len(s1Row)
		s2CharCount := len(s2Row)
		rowCharCount := maxInt(zeroOrPositive(-x)+s1CharCount, absInt(x)+s2CharCount)

		for i := 0; i < rowCharCount; i++ {

			var s1Char = ' '
			var s2Char = ' '

			if s1Index >= 0 && s1Index < s1CharCount {
				s1Char = s1Row[s1Index]
			}
			if s2Index >= 0 && s2Index < s2CharCount {
				s2Char = s2Row[s2Index]
			}
			if s2Char == ' ' && s1Char == ' ' {
				res.WriteString(" ")
			} else if s2Char == ' ' {
				res.WriteRune(s1Char)
			} else {
				res.WriteRune(s2Char)
			}
			s1Index++
			s2Index++
		}
		if r != rows-1 {
			res.WriteString("\n")
		}
	}
	return res.String()
}

// absInt returns the absolute value of an integer
func absInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

// zeroOrPositive returns zero if i is negative, otherwise i
func zeroOrPositive(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

// maxInt returns the maximum int between x and y
func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
