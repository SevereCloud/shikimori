package internal

import "container/list"

type fifo[T any] struct {
	list  *list.List
	limit int
}

func newFifo[T any]() *fifo[T] {
	return &fifo[T]{
		list:  list.New(),
		limit: 0,
	}
}

func (f *fifo[T]) checkLimit() {
	for f.Len() > f.limit {
		f.Pop()
	}
}

func (f *fifo[T]) Limit(limit int) {
	f.limit = limit

	if limit == 0 {
		return
	}

	f.checkLimit()
}

func (f *fifo[T]) Len() int {
	return f.list.Len()
}

func (f *fifo[T]) Push(element T) {
	f.list.PushBack(element)
	f.checkLimit()
}

func (f *fifo[T]) Front() *T {
	element := f.list.Front()
	if element == nil {
		return nil
	}

	value := element.Value.(T) //nolint:forcetypeassert

	return &value
}

func (f *fifo[T]) Pop() *T {
	element := f.list.Front()
	if element == nil {
		return nil
	}

	value := element.Value.(T) //nolint:forcetypeassert
	f.list.Remove(element)

	return &value
}
