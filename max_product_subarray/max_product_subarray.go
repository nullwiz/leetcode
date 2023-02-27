package main

import (
	"fmt"
)

// [2,3,4,5,6,-2,-3, -3, 2, -3,-5] 5 occ
// [1,-2,3,5,6,-7,1,2,-3,2,3,5,-3,-2]

// [1,2,-2,4,5,1-3] [-1,-2,3,4,5,6] [-2,-2,-3,-4,5,6] flip_c = even
// [-2,-4,-3,2,5] [1,2,-4,-2,5,-3] flic =  odd
// [2,3,4,5,5]

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	current_product := 1
	current_result := 0

	for i := 0; i < len(nums); i++ {
		current_product *= nums[i]
		current_result = max(current_product, current_result)
		if current_product == 0 {
			current_product = 1
		}
	}
	current_product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		current_product *= nums[i]
		current_result = max(current_product, current_result)
		if current_product == 0 {
			current_product = 1
		}
	}
	return current_result

}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func mod(i int) int {
	if i > 0 {
		return i
	}
	return i * (-1)
}

func main() {
	// runner
	a := []int{1, -1, 2, 3, -2, -20}
	fmt.Print(maxProduct(a))
}
