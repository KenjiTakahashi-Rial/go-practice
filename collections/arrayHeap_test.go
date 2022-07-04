package collections

import (
	"practice/test"
	"testing"
)

func TestArrayHeap(t *testing.T) {
	heap := MinHeap{}
	params := test.DefaultParams
	for i := params.RangeStart; i < params.RangeEnd; i++ {
		heap.Cap()
		if heap.Len() != i-params.RangeStart {
			t.Errorf("Heap shows length %d, should be %d\n", heap.Len(), i)
		}
		heap.Push(i)
	}

	length := params.RangeEnd - params.RangeStart
	for i := length - 1; i >= 0; i-- {
		peek, ok1 := heap.Peek()
		pop, ok2 := heap.Pop()
		if !ok1 || !ok2 || heap.Len() != i || peek != pop {
			t.Errorf("Peek ok? %t\nPop ok? %t\nLength (should be %d)? %d\nPeek val? %d\nPop val? %d\n", ok1, ok2, i, heap.Len(), peek, pop)
		}
	}
}
