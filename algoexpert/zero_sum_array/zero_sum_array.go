package main

import "fmt"

func ZeroSumSubarray(nums []int) bool {
	// Write your code
	hashSum := map[int]bool{}
	currSum := 0
	for _, val := range nums {
		currSum += val
		if val == 0 || currSum == 0 || hashSum[currSum] {
			return true
		}
		hashSum[currSum] = true
	}

	return false
}

func min(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	arr := []int{1, 2, 3, -2, -1}
	fmt.Println(ZeroSumSubarray(arr))
}
