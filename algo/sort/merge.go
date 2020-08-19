package sort

// MergeSort Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	left := MergeSort(slice[:mid])
	right := MergeSort(slice[mid:])
	return Merge(left, right)
}

// Merge Merges left and right slice into newly created slice
func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		lLeft, lRight := len(left)-1, len(right)-1
		if i > lLeft && j <= lRight {
			slice[k] = right[j]
			j++
		} else if j > lRight && i <= lLeft {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
