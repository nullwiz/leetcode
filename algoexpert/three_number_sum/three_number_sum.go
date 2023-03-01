package main

import "fmt"

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	res := [][]int{}
	currentSum := 0
	// 0 1 7
	// 0 2 7
	for i := 0; i <= len(array)-3; i++ {
		a := i + 1
		b := len(array) - 1
		for a < b {
			currentSum = array[i] + array[a] + array[b]
			if currentSum < target {
				a++
			}
			if currentSum > target {
				b--
			}
			if currentSum == target {
				res = append(res, []int{array[i], array[a], array[b]})
				b--
			}
		}
	}
	return res
}

// Returns a sorted slice of [n, i]
func getSortedSlice(arr []int, index int) []int {
	return QuickSort([]int{arr[index], arr[index+1], arr[index+2]})
}
func getSum(arr []int, start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	return sum
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
	arr := []int{12, 3, 1, 2, -6, 5, -8, 6}
	arr = QuickSort(arr)
	fmt.Println(ThreeNumberSum(arr, 0))
}
