package main

import "fmt"

type stack struct {
	left  int
	right int
}

func quicksortIterative(arr []int) {
	if len(arr) < 2 {
		return
	}

	s := []stack{{0, len(arr) - 1}}
	for len(s) > 0 {
		l, r := s[len(s)-1].left, s[len(s)-1].right
		s = s[:len(s)-1]

		if l >= r {
			continue
		}

		pivotIndex := partition(arr, l, r)
		s = append(s, stack{l, pivotIndex - 1})
		s = append(s, stack{pivotIndex + 1, r})
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{5, 2, 9, -3, 7, -4, 8, -1, 6}
	quicksortIterative(arr)
	fmt.Println(arr)
}
