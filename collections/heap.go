package collections

import (
	"container/heap"
	"practice/constraints"
)

// Private inner heap

type innerHeap []constraints.Ordered

func (h innerHeap) Len() int           { return len(h) }
func (h innerHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h innerHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *innerHeap) Push(x any) {
	*h = append(*h, x.(constraints.Ordered))
}

func (h *innerHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Public wrapper heap

type Heap[T constraints.Ordered] struct {
	inner *innerHeap
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	inner := make(innerHeap, 0)
	return &Heap[T]{&inner}
}

func (h *Heap[T]) Cap() int {
	return cap(*h.inner)
}

func (h *Heap[T]) Len() int {
	return h.inner.Len()
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.inner.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return (*h.inner)[0].(T), true
}

func (h *Heap[T]) Push(element T) {
	heap.Push(h.inner, element)
}

func (h *Heap[T]) Pop() (T, bool) {
	if h.inner.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return heap.Pop(h.inner).(T), true
}
