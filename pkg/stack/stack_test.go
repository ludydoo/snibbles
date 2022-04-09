package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(1)
	assert.Equal(t, 1, stack.Len())
}

func TestStack_Pop(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)

	var ok bool
	var v int

	v, ok = stack.Pop()
	assert.Equal(t, 2, v)
	assert.False(t, stack.IsEmpty())
	assert.True(t, ok)

	v, ok = stack.Pop()
	assert.Equal(t, 1, v)
	assert.True(t, stack.IsEmpty())
	assert.True(t, ok)

	v, ok = stack.Pop()
	assert.False(t, ok)
}
