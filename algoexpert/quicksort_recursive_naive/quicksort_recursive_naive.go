package main

import "fmt"

func quickSortNaive(arr []int) []int {
	// base case
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]
	// ups space complexity with tho additional lists. easiest case
	less := []int{}
	greater := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		}
		if arr[i] > pivot {
			greater = append(greater, arr[i])
		}
	}
	less = quickSortNaive(less)
	greater = quickSortNaive(greater)
	res := []int{}
	res = append(res, less...)
	res = append(res, pivot)
	res = append(res, greater...)
	return res
}

func main() {
	arr := []int{5, 2, 9, -3, 7, -4, 8, -1, 6}
	fmt.Println(quickSortNaive(arr))
}
