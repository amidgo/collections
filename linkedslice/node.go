package linkedslice

type nodes[T any] struct {
	items []*node[T]
}

func makeNodes[T any](initNode *node[T]) *nodes[T] {
	return &nodes[T]{
		items: []*node[T]{initNode},
	}
}

func (n *nodes[T]) InitNode() *node[T] {
	return n.items[0]
}

func (n *nodes[T]) LastNode() *node[T] {
	nodeCount := len(n.items) - 1
	return n.items[nodeCount]
}

func (n *nodes[T]) Node(index int) *node[T] {
	return n.items[index]
}

func (n *nodes[T]) AddNode(node *node[T]) {
	n.items = append(n.items, node)
}

type node[T any] struct {
	items  []T
	offset int
}

func makeNode[T any](offset, len, cap int) *node[T] {
	return &node[T]{
		offset: offset,
		items:  make([]T, len, cap),
	}
}

func (n *node[T]) Len() int {
	return len(n.items)
}

func (n *node[T]) Cap() int {
	return cap(n.items)
}

func (n *node[T]) OffsetIndex(index int) int {
	return index - n.offset
}

func (n *node[T]) SetItem(index int, item T) {
	n.items[index] = item
}

func (n *node[T]) Item(index int) T {
	return n.items[index]
}

func (n *node[T]) Add(item T) {
	n.items = append(n.items, item)
}
