package main

func threeSumHash(target int) (int, int, int) {
	target = 9
	arr := []int{1, 4, 5, 7, 1, 23, 8, 33}
	for i := 0; i < len(arr); i++ {
		targetSecondOrder := target - i
		if one, two := twoSumHash(targetSecondOrder); one != -1 && two != -1 {
			return one, two, i
		}
	}
	return -1, -1, -1
}
