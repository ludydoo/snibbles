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
