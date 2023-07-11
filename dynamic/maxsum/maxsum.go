package main

import "fmt"

func MaxSubsetSumNoAdjacent(array []int) int {
	// Write your code here.
	if len(array) < 1 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	maxSum := array
	maxSum[1] = max(array[0], array[1])
	for i := 2; i < len(array); i++ {
		maxSum[i] = max(maxSum[i-1], maxSum[i-2]+array[i])
	}
	fmt.Println(maxSum)
	return maxSum[len(array)-1]
}

func maxSubsetSumNoAdjacent(array []int) int {

	if len(array) < 1 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	second := array[0]
	first := max(array[0], array[1])
	for i := 2; i < len(array); i++ {
		current := max(first, second+array[i])
		second = first
		first = current
	}
	return first
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxSubsetSumNoAdjacent([]int{75, 105, 120, 75, 90, 135}))
	fmt.Println(maxSubsetSumNoAdjacent([]int{7, 10, 12, 7, 9, 14}))
	fmt.Println(maxSubsetSumNoAdjacent([]int{4, 3, 5, 200, 5, 3}))
}
