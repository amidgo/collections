package iterator

type Iterable[T any] interface {
	HasNext() bool
	Next() T
}

type Iterator[T any] struct {
	Iterable[T]
}

type IterateFunc[T any] func(int, T)

func (i Iterator[T]) Iterate(f IterateFunc[T]) {
	var index int
	for i.HasNext() {
		next := i.Next()
		f(index, next)
		index++
	}
}
