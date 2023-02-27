package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max_result := nums[0]
	current_sum := nums[0]

	for i := 1; i < len(nums); i++ {
		current_sum = max(nums[i], current_sum+nums[i])
		max_result = max(current_sum, max_result)
	}
	return max_result
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	// runner
	a := []int{-2, -1, 2, 4, 5, 6, -20, -30, 60}
	fmt.Print(maxSubArray(a))
}
