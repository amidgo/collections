package linkedslice_test

import (
	"testing"

	"github.com/amidgo/collections/linkedslice"
)

func addIntegersToSlice(x int) {
	intSlice := make([]int, 0, 10)
	for i := 0; i < x; i++ {
		intSlice = append(intSlice, i)
	}
}

func addPointersToSlice(x int) {
	intSlice := make([]*int, 0, 10)
	for i := 0; i < x; i++ {
		intSlice = append(intSlice, &i)
	}
}

func addIntegersToLinkedSlice(x int) {
	ls := linkedslice.MakeLinkedSlice[int](0, 10)
	for i := 0; i < x; i++ {
		ls.Add(i)
	}
}

func addPointersToLinkedSlice(x int) {
	ls := linkedslice.MakeLinkedSlice[*int](0, 10)
	for i := 0; i < x; i++ {
		ls.Add(&i)
	}
}

func BenchmarkSliceAdd_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToSlice(10000)
	}
}

func BenchmarkLinkedSliceAdd_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToLinkedSlice(10000)
	}
}

func BenchmarkSliceAdd_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToSlice(50000)
	}
}

func BenchmarkLinkedSliceAdd_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToLinkedSlice(50000)
	}
}

func BenchmarkSliceAdd_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToSlice(100000)
	}
}

func BenchmarkLinkedSliceAdd_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addIntegersToLinkedSlice(100000)
	}
}

func BenchmarkSliceAdd_Pointer_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToSlice(10000)
	}
}

func BenchmarkLinkedSliceAdd_Pointer_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToLinkedSlice(10000)
	}
}

func BenchmarkSliceAdd_Pointer_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToSlice(50000)
	}
}

func BenchmarkLinkedSliceAdd_Pointer_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToLinkedSlice(50000)
	}
}

func BenchmarkSliceAdd_Pointer_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToSlice(100000)
	}
}

func BenchmarkLinkedSliceAdd_Pointer_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPointersToLinkedSlice(100000)
	}
}
