package linkedslice

import (
	"github.com/amidgo/collections/iterator"
)

func (l *LinkedSlice[T]) Iterator() iterator.Iterator[T] {
	return iterator.Iterator[T]{
		Iterable: &LinkedSliceIteratable[T]{
			linkedSlice:          l,
			currentNodeItemIndex: -1,
		},
	}
}

type LinkedSliceIteratable[T any] struct {
	linkedSlice          *LinkedSlice[T]
	currentNodeIndex     int
	currentNodeItemIndex int
}

func (l *LinkedSliceIteratable[T]) HasNext() bool {
	linkedSliceLastNode := l.linkedSlice.nodes.LastNode()
	if l.currentNode() == linkedSliceLastNode {
		return l.currentNodeItemIndex != linkedSliceLastNode.Len()-1
	}
	return true
}

func (l *LinkedSliceIteratable[T]) Next() (item T) {
	l.next()
	return l.currentItem()
}

func (l *LinkedSliceIteratable[T]) next() {
	currentNode := l.currentNode()
	if currentNode.Len()-1 == l.currentNodeItemIndex {
		l.currentNodeIndex++
		l.currentNodeItemIndex = 0
	} else {
		l.currentNodeItemIndex++
	}
}

func (l *LinkedSliceIteratable[T]) currentNode() *node[T] {
	return l.linkedSlice.nodes.Node(l.currentNodeIndex)
}

func (l *LinkedSliceIteratable[T]) currentItem() T {
	return l.currentNode().Item(l.currentNodeItemIndex)
}
