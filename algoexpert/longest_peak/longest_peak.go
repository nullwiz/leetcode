package main

import "fmt"

func LongestPeak(array []int) int {
	currentPeak := []int{}
	longestPeak := []int{}
	mustIncrease := true
	mustDecrease := false
	for i := 1; i < len(array); i++ {
		curr := longestPeak[len(longestPeak)-1]
		if array[i] >= curr && mustIncrease {
			currentPeak = append(currentPeak, array[i])
		}
		if array[i] < curr {
			mustDecrease, mustIncrease = true, false
			currentPeak = append(currentPeak, array[i])
		}
	}
	return 1
}

func main() {
	fmt.Println("vim-go")
}
