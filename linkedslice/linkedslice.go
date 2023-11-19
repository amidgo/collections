package linkedslice

import (
	"fmt"
	"math"
)

type LinkedSlice[T any] struct {
	len, cap int
	nodes    nodes[T]
}

func MakeLinkedSlice[T any](len, cap int) *LinkedSlice[T] {
	initNode := makeNode[T](0, len, cap)
	return &LinkedSlice[T]{
		len:   len,
		cap:   cap,
		nodes: makeNodes[T](initNode),
	}
}

func (l *LinkedSlice[T]) Len() int {
	return l.len
}

func (l *LinkedSlice[T]) Cap() int {
	return l.cap
}

func (l *LinkedSlice[T]) Swap(firstIndex, secondIndex int) {
	firstItem := l.Item(firstIndex)
	secondItem := l.Item(secondIndex)
	l.SetItem(secondIndex, firstItem)
	l.SetItem(firstIndex, secondItem)
}

func (l *LinkedSlice[T]) SetItem(index int, item T) {
	if index >= l.len {
		l.panicIndexOutOfRange(index)
	}
	node := l.searchNode(index)
	nodeItemIndex := node.OffsetIndex(index)
	node.SetItem(nodeItemIndex, item)
}

func (l *LinkedSlice[T]) Item(index int) T {
	if index >= l.len {
		l.panicIndexOutOfRange(index)
	}
	node := l.searchNode(index)
	nodeItemIndex := node.OffsetIndex(index)
	return node.Item(nodeItemIndex)
}

func (l *LinkedSlice[T]) panicIndexOutOfRange(index int) {
	message := fmt.Sprintf("index %d out of range [0:%d]", index, l.len-1)
	panic(message)
}

// illustration for init node capacity = 10
//
//	0	   1      2      3        4	       5
//
// 0..9 10..29 30..69 70..149 150..309 310..629
func (l *LinkedSlice[T]) searchNode(itemIndex int) *node[T] {
	initNode := l.nodes.InitNode()
	offset := itemIndex / initNode.Cap()
	offsetLog2 := math.Log2(float64(offset + 1))
	nodeIndex := int(offsetLog2)
	return l.nodes.Node(nodeIndex)
}

func (l *LinkedSlice[T]) AddItems(items ...T) {
	for i := range items {
		l.Add(items[i])
	}
}

func (l *LinkedSlice[T]) Add(item T) {
	l.addNodeIfOverflow()
	l.add(item)
}

func (l *LinkedSlice[T]) addNodeIfOverflow() {
	if l.overflow() {
		l.addNode()
	}
}

func (l *LinkedSlice[T]) overflow() bool {
	lastNode := l.nodes.LastNode()
	return lastNode.Cap() == lastNode.Len()
}

func (l *LinkedSlice[T]) addNode() {
	lastNode := l.nodes.LastNode()
	nodeCap := lastNode.Cap() * 2
	node := makeNode[T](l.cap, 0, nodeCap)
	l.nodes.AddNode(node)
	l.cap += nodeCap
}

func (l *LinkedSlice[T]) add(item T) {
	lastNode := l.nodes.LastNode()
	lastNode.Add(item)
	l.len++
}
