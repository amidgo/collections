package linkedslice_test

import (
	"testing"

	"github.com/amidgo/collections/linkedslice"
)

func BenchmarkSliceItemByIndex(b *testing.B) {
	items := GenerateIntegerSlice(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = items[i]
	}
}

func BenchmarkMapItemByKey(b *testing.B) {
	items := GenerateIntegerMap(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = items[i]
	}
}

func BenchmarkLinkedSliceItemByIndex(b *testing.B) {
	ls := linkedslice.MakeLinkedSlice[int](0, 10)
	items := GenerateIntegerSlice(b.N)
	ls.AddItems(items...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ls.Item(i)
	}
}
