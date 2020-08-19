package sort

import (
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// LaunchInsertionSort represents ins sort example
func LaunchInsertionSort() []int {
	// slice := generateSlice(20)
	slice := []int{5, 88, 75, 22, 7, 1, 2, 4}
	spew.Dump(" Unsorted ", slice)
	// fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionsort(slice)
	spew.Dump(" Sorted ", slice)
	// fmt.Println("\n--- Sorted ---\n\n", slice)
	return slice
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
