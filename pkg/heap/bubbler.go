package heap

import "golang.org/x/exp/constraints"

func minBubbleUp[V constraints.Ordered](h *heap[V], i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[parent] > (h.elements[i]) {
			h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
			i = parent
		} else {
			break
		}
	}
}

func minBubbleDown[V constraints.Ordered](h *heap[V], i int) {
	for i < h.size/2 {
		left := 2*i + 1
		right := 2*i + 2
		smallest := left
		if right < h.size && h.elements[right] < (h.elements[left]) {
			smallest = right
		}
		if h.elements[i] < (h.elements[smallest]) {
			break
		}
		h.elements[i], h.elements[smallest] = h.elements[smallest], h.elements[i]
		i = smallest
	}
}

func maxBubbleUp[V constraints.Ordered](h *heap[V], i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[parent] < h.elements[i] {
			h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
			i = parent
		} else {
			break
		}
	}
}

func maxBubbleDown[V constraints.Ordered](h *heap[V], i int) {
	if i == h.size-1 {
		return
	}
	left := 2*i + 1
	right := 2*i + 2
	if left >= h.size {
		return
	}
	if right >= h.size {
		if h.elements[i] < h.elements[left] {
			h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
			h.bubbleDown(h, left)
		}
		return
	}
	if h.elements[left] < h.elements[right] {
		if h.elements[i] < h.elements[right] {
			h.elements[i], h.elements[right] = h.elements[right], h.elements[i]
			h.bubbleDown(h, right)
		}
		return
	}
	if h.elements[i] < h.elements[left] {
		h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
		h.bubbleDown(h, left)
	}
}
