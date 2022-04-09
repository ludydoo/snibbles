package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {
	ll := LinkedList[int]{}
	idx := ll.Add(1)
	expected := LinkedList[int]{Head: &Node[int]{Value: 1, Next: nil}}
	assert.Equal(t, expected, ll)
	assert.Equal(t, 0, idx)
}

func TestLinkedList_AddSecond(t *testing.T) {
	ll := LinkedList[int]{}
	idx1 := ll.Add(1)
	idx2 := ll.Add(2)
	expected := LinkedList[int]{Head: &Node[int]{Value: 1, Next: &Node[int]{Value: 2, Next: nil}}}
	assert.Equal(t, expected, ll)
	assert.Equal(t, 0, idx1)
	assert.Equal(t, 1, idx2)
}

func TestLinkedList_RemoveAt(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.RemoveAt(1)
	expected := LinkedList[int]{Head: &Node[int]{Value: 1}}
	assert.Equal(t, expected, ll)
}

func TestLinkedList_RemoveAt_Head(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.RemoveAt(0)
	expected := LinkedList[int]{Head: &Node[int]{Value: 2}}
	assert.Equal(t, expected, ll)
}

func TestLinkedList_Remove(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	idx := ll.Remove(1)
	assert.True(t, ll.IsEmpty())
	assert.Equal(t, 0, idx)
}

func TestLinkedList_RemoveNotExisting(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	idx := ll.Remove(2)
	assert.False(t, ll.IsEmpty())
	assert.Equal(t, -1, idx)
}

func TestLinkedList_ElementAt(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	actual := ll.ElementAt(1)
	assert.Equal(t, &Node[int]{Value: 2}, actual)
}

func TestLinkedList_ElementAt_Head(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	actual := ll.ElementAt(0)
	assert.Equal(t, &Node[int]{Value: 1}, actual)
}

func TestLinkedList_Insert_Head(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Insert(0, 0)
	expected := LinkedList[int]{Head: &Node[int]{Value: 0, Next: &Node[int]{Value: 1, Next: &Node[int]{Value: 2, Next: nil}}}}
	assert.Equal(t, expected, ll)
}

func TestLinkedList_Insert_Middle(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Insert(0, 1)
	expected := LinkedList[int]{Head: &Node[int]{Value: 1, Next: &Node[int]{Value: 0, Next: &Node[int]{Value: 2, Next: nil}}}}
	assert.Equal(t, expected, ll)
}

func TestLinkedList_Insert_End(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Insert(0, 2)
	expected := LinkedList[int]{Head: &Node[int]{Value: 1, Next: &Node[int]{Value: 2, Next: &Node[int]{Value: 0, Next: nil}}}}
	assert.Equal(t, expected, ll)
}

func TestLinkedList_IsEmpty(t *testing.T) {
	ll := LinkedList[int]{}
	actual := ll.IsEmpty()
	assert.True(t, actual)
}

func TestLinkedList_IsEmpty_Not(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	actual := ll.IsEmpty()
	assert.False(t, actual)
}
