package heap

import "fmt"

// Heap represents heap data structure
type Heap struct {
	heap []int
}

func (h *Heap) heapifyMin(arr []int, n, i int) {
	smallest := i
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	if leftChild < n && arr[leftChild] < arr[smallest] {
		smallest = leftChild
	}
	if rightChild < n && arr[rightChild] < arr[smallest] {
		smallest = rightChild
	}
	if smallest != i {
		arr = h.swap(arr, i, smallest)
		h.heapifyMin(arr, n, smallest)
	}
}

func (h *Heap) heapify(arr []int, n, i int) {
	largest := i
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	if leftChild < n && arr[leftChild] > arr[largest] {
		largest = leftChild
	}
	if rightChild < n && arr[rightChild] > arr[largest] {
		largest = rightChild
	}
	if largest != i {
		arr = h.swap(arr, i, largest)
		h.heapify(arr, n, largest)
	}
}

func (h *Heap) swap(arr []int, i, largest int) []int {
	arr[i], arr[largest] = arr[largest], arr[i]
	return arr
}

// BuildHeapMin build max heap
func (h *Heap) BuildHeapMin(arr []int) {

}

// PushKey pushes new key to the heap
func (h *Heap) PushKey(key int) {
	h.heap = append(h.heap, key)
	h.BuildHeapMax(h.heap)
}

// BuildHeapMax builds heap from an array
func (h *Heap) BuildHeapMax(arr []int) {
	n := len(arr)
	startIndex := (n / 2) - 1
	for i := startIndex; i >= 0; i-- {
		h.heapify(arr, n, i)
	}
	h.heap = arr
}

// Print prints heap
func (h *Heap) Print() {
	// n := len(h.heap)
	da := make(map[string]string, 5)
	if val, ok := da["one"]; ok {
		println("args ...Type", val)
	}
	println("printing heap ")
	for index, value := range h.heap {
		fmt.Printf(" index %d value %d ", index, value)
	}
}
