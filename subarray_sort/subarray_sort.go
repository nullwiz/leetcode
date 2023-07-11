package main

import "fmt"

func SubarraySort(array []int) []int {
	// Loop from the left, find the indices
	ret := []int{-1, -1}
	// from the left
	inSub := false
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] && !inSub {
			inSub = true
			ret[0] = i
		}
		if array[i] < array[i+1] && inSub {
			inSub = false
			ret[1] = i
		}
		if array[i] > array[i+1] && !inSub {
		}
	}
	return ret
}

// [1,2,-2,3,4,3,1,2]
func main() {
	fmt.Println("vim-go")
}
