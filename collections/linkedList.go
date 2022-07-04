package collections

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type LinkedList[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	tail := &node[T]{}
	head := &node[T]{}
	tail.prev = head
	head.next = tail

	return &LinkedList[T]{0, head, tail}
}

func (l *LinkedList[T]) remove(n *node[T]) T {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
	l.length--
	return n.value
}

func (l *LinkedList[T]) Len() int {
	return l.length
}

func (l *LinkedList[T]) Front() (T, bool) {
	if l.length == 0 {
		var value T
		return value, false
	}
	return l.head.next.value, true
}

func (l *LinkedList[T]) Back() (T, bool) {
	if l.length == 0 {
		var value T
		return value, false
	}
	return l.tail.prev.value, true
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.length == 0 {
		var value T
		return value, false
	}

	n := l.head.next
	l.head.next = n.next
	n.next.prev = l.head
	return n.value, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	if l.length == 0 {
		var value T
		return value, false
	}

	n := l.tail.prev
	l.tail.prev = n.prev
	n.prev.next = l.tail
	return n.value, true
}

func (l *LinkedList[T]) PushFront(value T) {
	next := l.head.next
	n := &node[T]{value, l.head, next}
	l.head.next = n
	next.prev = n
	l.length++
}

func (l *LinkedList[T]) PushBack(value T) {
	prev := l.tail.prev
	n := &node[T]{value, prev, l.tail}
	prev.next = n
	l.tail.prev = n
	l.length++
}

func (l *LinkedList[T]) Slice() []T {
	slice := make([]T, l.length)
	curr := l.head.next
	for i := 0; curr != l.tail; i, curr = i+1, curr.next {
		slice[i] = curr.value
	}
	return slice
}

func RemoveFirst[T comparable](l *LinkedList[T], value T) (T, bool) {
	curr := l.head.next
	for curr != l.tail {
		if curr.value == value {
			return l.remove(curr), true
		}
		curr = curr.next
	}

	var zeroed T
	return zeroed, false
}

func RemoveLast[T comparable](l *LinkedList[T], value T) (T, bool) {
	curr := l.tail.prev
	for curr != l.head {
		if curr.value == value {
			return l.remove(curr), true
		}
		curr = curr.prev
	}

	var zeroed T
	return zeroed, false
}
