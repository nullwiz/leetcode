package main

import "fmt"

func LongestPeak(array []int) int {
	m := make(map[int][2]int)

	if len(array) < 2 {
		return 0
	}
	// O(n)
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			fmt.Printf("Peak found at %d with value %d\n", i, array[i])
			m[i] = [2]int{array[i], 0}
		}
	}
	
	fmt.Println(m)
	return 0
}
func main() {
	arr := []int{5, 4, 3, 2, 1, 2, 10, 12}
	fmt.Println(LongestPeak(arr))
	arr2 := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, 3, 2, 3}
	fmt.Println(LongestPeak(arr2))
}
