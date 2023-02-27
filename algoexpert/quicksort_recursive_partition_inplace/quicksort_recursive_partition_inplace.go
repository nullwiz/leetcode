package main

import "fmt"

func QuickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose last element as pivot
	pivot := arr[high]
	// Initialize index of smaller element
	i := low - 1
	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Swap arr[i+1] and arr[high] (pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{5, 2, 9, -3, 7, -4, 8, -1, 6}
	QuickSort(arr)
	fmt.Println(arr)
}
