package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(1)
	assert.Equal(t, 1, q.Len())
}

func TestQueue_Dequeue(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(1)
	v := q.Dequeue()
	assert.Equal(t, 1, v)
	assert.True(t, q.IsEmpty())
}

func BenchmarkQueue_Enqueue(b *testing.B) {
	q := &Queue[int]{}
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkQueue_Dequeue(b *testing.B) {
	q := &Queue[int]{}
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
