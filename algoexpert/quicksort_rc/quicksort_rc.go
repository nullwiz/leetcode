package main

import "math/rand"

func quicksort(arr []int) []int {
	// Base case
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	// Move the pivot to the right
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Pile elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}

func main(){
	arr := []int{5,4,-1,3,2,1}
	quicksort(arr)
}