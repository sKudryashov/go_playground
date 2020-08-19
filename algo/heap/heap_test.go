package heap_test

import (
	"testing"

	"github.com/sKudryashov/algo/heap"
)

func TestHeap(t *testing.T) {
	data := []int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17}
	h := &heap.Heap{}
	h.BuildHeapMax(data)
	h.Print()
	h.PushKey(22)
	h.Print()
}

func TestNativeHeap(t *testing.T) {
	heap.LaunchHeap()
}
