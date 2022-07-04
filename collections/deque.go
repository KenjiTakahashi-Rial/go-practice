package collections

import "container/list"

type Deque[T any] struct {
	list *list.List
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{list.New()}
}

func (l *Deque[T]) Len() int {
	return l.list.Len()
}

func (l *Deque[T]) Clear() {
	l.list = l.list.Init()
}

func (l *Deque[T]) Back() (T, bool) {
	if l.list.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return l.list.Back().Value.(T), true
}

func (l *Deque[T]) Front() (T, bool) {
	if l.list.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return l.list.Front().Value.(T), true
}

func (l *Deque[T]) PopBack() (T, bool) {
	if l.list.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return l.list.Remove(l.list.Back()).(T), true
}

func (l *Deque[T]) PopFront() (T, bool) {
	if l.list.Len() == 0 {
		var zeroed T
		return zeroed, false
	}
	return l.list.Remove(l.list.Front()).(T), true
}

func (l *Deque[T]) PushBack(v T) {
	l.list.PushBack(v)
}

func (l *Deque[T]) PushFront(v T) {
	l.list.PushFront(v)
}

func (l *Deque[T]) ConcatBack(other *Deque[T]) {
	l.list.PushBackList(other.list)
}

func (l *Deque[T]) ConcatFront(other *Deque[T]) {
	l.list.PushFrontList(other.list)
}
