package main

import "fmt"

func IsMonotonic(array []int) bool {
	// Write your code here.
	if len(array) < 2 {
		return true
	}
	helperarray := make([]int, len(array))
	copy(helperarray, array)
	j := 0
	if array[0] < array[len(array)-1] {
		helperarray[0] = 1
	} else {
		helperarray[0] = 0
	}
	sum := 0
	for i := 1; i < len(array); i++ {
		pre := array[j]
		if array[i] > pre {
			helperarray[i] = 1
			sum++
		}
		if array[i] < pre {
			helperarray[i] = 0
		}
		if array[i] == pre {
			helperarray[i] = helperarray[j]
			if helperarray[j] == 1 {
				sum++
			}
		}
		j++
	}
	if sum == len(array)-1 || sum == 0 {
		return true
	} else {
		return false
	}
}

func IsMonotonicAndNotRetarded(array []int) bool {
	// Write your code here.
	if len(array) < 2 {
		return true
	}
	isNonDecreasing := true
	isNonIncreasing := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecreasing = false
		}
		if array[i] > array[i-1] {
			isNonIncreasing = false
		}
	}
	return isNonDecreasing || isNonIncreasing
}
func main() {
	arr := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	fmt.Println(IsMonotonic(arr))
}
