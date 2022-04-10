package kruskal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKruskal(t *testing.T) {

	graph := []Edge{
		{1, 2, 2},
		{1, 4, 1},
		{1, 5, 4},
		{5, 4, 9},
		{4, 3, 5},
		{4, 2, 3},
		{2, 3, 3},
		{2, 6, 7},
		{6, 3, 8},
	}

	expected := []Edge{
		{1, 4, 1},
		{1, 2, 2},
		{2, 3, 3},
		{1, 5, 4},
		{2, 6, 7},
	}

	print("Input: \n")
	printGraph(graph)

	result := Kruskal(graph)

	print("Output: \n")
	printGraph(result)

	assert.True(t, graphsEqual(result, expected))
}
