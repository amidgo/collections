package linkedslice

import "github.com/amidgo/collections/iterator"

func (l *LinkedSlice[T]) Iterator() iterator.Iterator[T] {
	return iterator.Iterator[T]{
		Iterable: &LinkedSliceIteratable[T]{
			linkedSlice: l,
		},
	}
}

type LinkedSliceIteratable[T any] struct {
	linkedSlice *LinkedSlice[T]
	index       int
}

func (l *LinkedSliceIteratable[T]) HasNext() bool {
	return l.index < l.linkedSlice.Len()-1
}

func (l *LinkedSliceIteratable[T]) Next() (item T) {
	item = l.linkedSlice.Item(l.index)
	l.index++
	return
}
