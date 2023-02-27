package main

import "fmt"

type stack struct {
	left  int
	right int
}

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	// |x| - |y| = 0
	var curDiff int
	a, b := 0, 0
	c, d := 0, 0
	le1 := len(array1) - 1
	le2 := len(array2) - 1
	//O(nlog(n)) + O(mlog(m))
	quicksortIterative(array1)
	quicksortIterative(array2)
	//O(n+m)
	for a < le1 && b < le2 {
		val1 := abs(array1[a])
		val2 := abs(array2[b])
		if val1 == val2 {
			break
		}
		if curDiff == 0 {
			curDiff = val1 - val2
		}
		if abs(curDiff) > abs(val1-val2) {
			curDiff = val1 - val2
			c, d = a, b
		}
		if val1 < val2 {
			a++
		}
		if val2 < val1 {
			b++
		}
	}
	for a < le1 {
		val1 := abs(array1[a])
		val2 := abs(array2[b])
		if abs(curDiff) > abs(val1-val2) {
			curDiff = val1 - val2
			c, d = a, b
		}
		a++
	}
	for b < le2 {
		val1 := abs(array1[a])
		val2 := abs(array2[b])
		if abs(curDiff) > abs(val1-val2) {
			curDiff = val1 - val2
			c, d = a, b
		}
		b++
	}
	return []int{array1[c], array2[d]}
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

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}

func main() {
	arrOne := []int{-1, 5, 10, 20, 3}
	arrTwo := []int{26, 134, 135, 15, 17}
	answ := SmallestDifference(arrOne, arrTwo)
	fmt.Println(answ)
}
