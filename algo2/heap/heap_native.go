package heap

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
// []*Item

type Item struct {
	Index int
}
type IntHeap []*Item

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Index < h[j].Index }
func (h IntHeap) Swap(i, j int)      { h[i].Index, h[j].Index = h[j].Index, h[i].Index }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// LaunchHeap
func LaunchHeap() {
	h := &IntHeap{&Item{Index: 2}, &Item{Index: 1}, &Item{Index: 5}}
	heap.Init(h)
	heap.Push(h, &Item{Index: 3})
	heap.Push(h, &Item{Index: 8})
	fmt.Printf("minimum: %d\n", (*h)[0].Index)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
