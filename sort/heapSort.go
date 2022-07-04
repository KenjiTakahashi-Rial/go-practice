package sort

import (
	"practice/collections"
	"practice/constraints"
)

func HeapSort(arr []int) {
	heap := collections.NewHeap[constraints.OrderedInt]()

	for _, n := range arr {
		heap.Push(constraints.OrderedInt(n))
	}

	for i := 0; i < len(arr); i++ {
		pop, _ := heap.Pop()
		arr[i] = int(pop)
	}
}
