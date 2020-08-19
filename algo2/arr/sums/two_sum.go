package main

func main() {
	var one, two int
	one, two = twoSumHash(9)
	println("one, two ", one, two)
	one, two = twoSumPointer(9)
	println("one, two ", one, two)
}

func twoSumPointer(target int) (int, int) {
	arrSorted := []int{2, 7, 11, 15, 16, 18}
	var low, hi int
	hi = len(arrSorted) - 1
	for low <= hi {
		t := arrSorted[low] + arrSorted[hi]
		if t == target {
			return low, hi
		} else if target < hi {
			hi--
		} else {
			low++
		}
	}
	return hi, low
}

func twoSumHash(target int) (int, int) {
	arr := []int{2, 7, 11, 15, 7, 12}
	accumulator := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		accumulator[arr[i]] = i
	}
	for i := 0; i < len(arr); i++ {
		diff := target - arr[i]
		if _, ok := accumulator[diff]; ok {
			return i, diff
		}
	}
	return -1, -1
}
