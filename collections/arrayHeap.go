package collections

type MinHeap struct {
	elements []int
}

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func (heap *MinHeap) swap(i, j int) {
	e := heap.elements
	e[i], e[j] = e[j], e[i]
}

func (heap *MinHeap) heapifyUp() {
	i := len(heap.elements) - 1
	for {
		p := parent(i)
		if i == 0 || heap.elements[i] >= heap.elements[p] {
			break
		}

		heap.swap(i, p)
		i = p
	}
}

func (heap *MinHeap) heapifyDown() {
	i := 0
	for {
		l, r := left(i), right(i)

		var lower int
		if r >= len(heap.elements) || heap.elements[l] < heap.elements[r] {
			lower = l
		} else {
			lower = r
		}

		if l >= len(heap.elements) || heap.elements[i] <= heap.elements[lower] {
			break
		}

		heap.swap(i, lower)
		i = lower
	}
}

func (heap *MinHeap) Len() int {
	return len(heap.elements)
}

func (heap *MinHeap) Cap() int {
	return cap(heap.elements)
}

func (heap *MinHeap) Peek() (int, bool) {
	if len(heap.elements) == 0 {
		return 0, false
	}
	return heap.elements[0], true
}

func (heap *MinHeap) Push(element int) {
	heap.elements = append(heap.elements, element)
	heap.heapifyUp()
}

func (heap *MinHeap) Pop() (int, bool) {
	if len(heap.elements) == 0 {
		return 0, false
	}
	popped := heap.elements[0]
	heap.elements[0] = heap.elements[len(heap.elements)-1]
	heap.elements = heap.elements[:len(heap.elements)-1]
	heap.heapifyDown()
	return popped, true
}
