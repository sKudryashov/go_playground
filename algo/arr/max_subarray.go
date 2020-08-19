package arr

import "fmt"

func LaunchMaxSubArr() {
	subar := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArr(subar)
	fmt.Printf("max %d", max)
}

func maxSubArr(nums []int) int {
	var currentSum int
	var maxSum int
	//[-2,1,-3,4,-1,2,1,-5,4]
	for i := 0; i < len(nums); i++ {
		currentSummProgress := currentSum + nums[i]
		currentSum = max(nums[i], currentSummProgress)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(one, two int) int {
	if one > two {
		return one
	}
	return two
}
