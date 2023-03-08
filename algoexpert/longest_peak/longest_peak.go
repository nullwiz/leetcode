package main

import "fmt"

func LongestPeak(array []int) int {
	currentPeak := []int{array[0]}
	longestPeak := []int{array[0]}
	mustIncrease := true
	for i := 1; i < len(array); i++ {
		curr := longestPeak[len(longestPeak)-1]

		if array[i] >= curr && mustIncrease {
			currentPeak = append(currentPeak, array[i])
		}
		if array[i] < curr {
			if array[i-1] == maxOf(currentPeak) {
				mustIncrease = false
				currentPeak = append(currentPeak, array[i])
			} else {
				mustIncrease = true
				currentPeak = []int{array[i]}
			}
		}

		if len(currentPeak) > len(longestPeak) {
			longestPeak = currentPeak
		}
	}

	return len(longestPeak)
}

func maxOf(array []int) int {
	currMax := array[0]
	for _, v := range array[1:] {
		if v >= currMax {
			currMax = v
		}
	}
	return currMax
}
func main() {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	fmt.Println(LongestPeak(arr))
}
