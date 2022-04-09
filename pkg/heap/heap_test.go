package heap

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"sort"
	"testing"
)

func TestHeap_Insert(t *testing.T) {

	const capacity = 16

	maxHeap := Max[int](capacity)
	minHeap := Min[int](capacity)

	ints := randomInts(capacity)

	ascending := sortInts(ints)
	descending := reverseInts(ascending)

	for _, v := range ints {
		maxHeap.Push(v)
		minHeap.Push(v)
	}

	for i := range ints {
		maxHeapElem, _ := maxHeap.Pop()
		minHeapElem, _ := minHeap.Pop()

		expectedMaxHeapElem := descending[i]
		expectedMinHeapElem := ascending[i]

		assert.Equal(t, expectedMaxHeapElem, maxHeapElem)
		assert.Equal(t, expectedMinHeapElem, minHeapElem)
	}

}

func sortInts(ints []int) []int {
	intCopy := make([]int, len(ints))
	copy(intCopy, ints)
	sort.Ints(intCopy)
	return intCopy
}

func reverseInts(ints []int) []int {
	c := make([]int, len(ints))
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = ints[j], ints[i]
	}
	return c
}

func randomInts(n int) []int {
	ints := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}
	rand.Shuffle(n, func(i, j int) {
		ints[i], ints[j] = ints[j], ints[i]
	})
	return ints
}
