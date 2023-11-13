package linkedslice_test

import (
	"math/rand"
	"testing"

	"github.com/amidgo/collections/linkedslice"
	"github.com/stretchr/testify/assert"
)

func GenerateIntegerSlice(size int) []int {
	intSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		intSlice = append(intSlice, rand.Int())
	}
	return intSlice
}

func GenerateIntegerMap(size int) map[int]struct{} {
	intMap := make(map[int]struct{}, size)
	for i := 0; i < size; i++ {
		intMap[i] = struct{}{}
	}
	return intMap
}

func Test_LinkedSlice_Add_And_Item(t *testing.T) {
	cases := []struct {
		initCap    int
		itemsCount int
	}{
		{
			initCap:    14,
			itemsCount: 10000,
		},
		{
			initCap:    1,
			itemsCount: 12312,
		},
	}

	for _, cs := range cases {
		ls := linkedslice.MakeLinkedSlice[int](0, cs.initCap)
		intSlice := GenerateIntegerSlice(cs.itemsCount)
		ls.AddItems(intSlice...)

		for i := range intSlice {
			assert.Equal(t, intSlice[i], ls.Item(i))
		}
	}
}

func Test_LinkedSlice_Len_Cap(t *testing.T) {
	cases := []struct {
		initCap     int
		itemsCount  int
		expectedCap int
	}{
		{
			initCap:     1,
			itemsCount:  96,
			expectedCap: 127,
		},
		{
			initCap:     3,
			itemsCount:  56,
			expectedCap: 93,
		},
	}

	for _, cs := range cases {
		ls := linkedslice.MakeLinkedSlice[int](0, cs.initCap)
		intSlice := GenerateIntegerSlice(cs.itemsCount)
		ls.AddItems(intSlice...)

		assert.Equal(t, cs.itemsCount, ls.Len())
		assert.Equal(t, cs.expectedCap, ls.Cap())
	}
}

func Test_LinkedSlice_SetItem(t *testing.T) {
	cases := []struct {
		initCap    int
		itemsCount int

		setItemIndex int
		setItem      int
	}{
		{
			initCap:      8,
			itemsCount:   1000,
			setItemIndex: 631,
			setItem:      1923812381238,
		},
	}

	for _, cs := range cases {
		ls := linkedslice.MakeLinkedSlice[int](0, cs.initCap)
		intSlice := GenerateIntegerSlice(cs.itemsCount)
		ls.AddItems(intSlice...)

		assert.NotEqual(t, intSlice[cs.setItemIndex], cs.setItem)

		ls.SetItem(cs.setItemIndex, cs.setItem)

		assert.Equal(t, cs.setItem, ls.Item(cs.setItemIndex))
	}
}

func Test_LinkedSlice_Swap(t *testing.T) {
	cases := []struct {
		initCap         int
		itemsCount      int
		firstSwapIndex  int
		secondSwapIndex int
	}{
		{
			initCap:         10,
			itemsCount:      10001,
			firstSwapIndex:  231,
			secondSwapIndex: 0,
		},
	}

	for _, cs := range cases {
		ls := linkedslice.MakeLinkedSlice[int](0, cs.initCap)
		intSlice := GenerateIntegerSlice(cs.itemsCount)
		ls.AddItems(intSlice...)

		ls.Swap(cs.firstSwapIndex, cs.secondSwapIndex)

		assert.Equal(t, ls.Item(cs.firstSwapIndex), intSlice[cs.secondSwapIndex])
		assert.Equal(t, ls.Item(cs.secondSwapIndex), intSlice[cs.firstSwapIndex])
	}
}

func Test_LinkedSlice_OutOfRange_Item(t *testing.T) {
	defer func() {
		message := recover()
		assert.Equal(t, message, "index 100 out of range [0:-1]")
	}()
	ls := linkedslice.MakeLinkedSlice[int](0, 10)
	ls.Item(100)
}

func Test_LinkedSlice_OutOfRange_SetItem(t *testing.T) {
	defer func() {
		message := recover()
		assert.Equal(t, message, "index 100 out of range [0:-1]")
	}()
	ls := linkedslice.MakeLinkedSlice[int](0, 10)
	ls.SetItem(100, rand.Int())
}

func Test_LinkedSlice_Iterator(t *testing.T) {
	cases := []struct {
		initCap    int
		itemsCount int
	}{
		{
			initCap:    10,
			itemsCount: 102321,
		},
	}

	for _, cs := range cases {
		intSlice := GenerateIntegerSlice(cs.itemsCount)
		ls := linkedslice.MakeLinkedSlice[int](0, 10)
		ls.AddItems(intSlice...)
		var iterableFunc = func(index, item int) {
			assert.Equal(t, intSlice[index], item)
		}
		ls.Iterator().Iterate(iterableFunc)
	}
}
