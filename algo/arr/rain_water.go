package arr 

func Trap(height []int) int {
	var rightMax, leftMax, right, left, sum int
	left = 0 
	right = len(height) - 1
	for left < right {
		// advance left (how to undertdand when to advance left?)
		// advance right (how to undertdand when to advance right?)
		if height[left] < height[right] {
			if leftMax > height[left] {
				sum = sum + leftMax - height[left]
			} else if leftMax <= height[left] {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right]  {
				sum = sum + rightMax - height[right]
			} else if rightMax <= height[right] {
				rightMax = height[right] 
			}
			right--
		}
	}
	return sum
}