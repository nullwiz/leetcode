package main

import "fmt"

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	if len(coins) == 0 {
		return 1
	}
	return 0
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}
	ret := []int{}

	for _, v := range arr[1:] {
		if v > pivot {
			greater = append(greater, v)
		}
		if v <= pivot {
			less = append(less, v)
		}
	}
	less = QuickSort(less)
	greater = QuickSort(greater)
	ret = append(ret, less...)
	ret = append(ret, pivot)
	ret = append(ret, greater...)
	return ret
}

func main() {
	arr := []int{1, -1, 1, -2, -4, -6, -8, -9, 10, 20}
	fmt.Println(QuickSort(arr))
}
