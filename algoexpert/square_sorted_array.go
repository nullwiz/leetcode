package main

import "fmt"

func SortedSquaredArray(array []int) []int {
	length := len(array) - 1
	// Base case
	arr := make([]int, length+1)
	small := 0
	big := length
	for small <= big {
		big_int := array[big]
		smol_int := array[small]
		if abs(big_int) > abs(smol_int) {
			arr[length] = big_int * big_int
			big--
		}
		if abs(big_int) < abs(smol_int) {
			arr[length] = smol_int * smol_int
			small++

		}
		if abs(big_int) == abs(smol_int) {
			arr[length] = smol_int * smol_int
			if big != small {
				arr[length-1] = smol_int * smol_int
			}
			small++
			big--
			length = length - 2
			continue
		}
		length--
	}
	return arr
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -1 * x
	}
}

func main() {
	a := []int{-7, -3, 1, 9, 22, 30}
	fmt.Println(SortedSquaredArray(a))
}
