package main

import (
	"fmt"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	// Write your code here.
	sort.Ints(redShirtSpeeds)
	if fastest {
		sort.Slice(blueShirtSpeeds, func(i, j int) bool {
			return blueShirtSpeeds[i] > blueShirtSpeeds[j]
		})
	} else {
		sort.Slice(blueShirtSpeeds, func(i, j int) bool {
			return blueShirtSpeeds[i] < blueShirtSpeeds[j]
		})

	}
	sum := 0
	fmt.Println(redShirtSpeeds)
	fmt.Println(blueShirtSpeeds)
	for i := 0; i < len(redShirtSpeeds); i++ {
		fmt.Printf("pairs [%d,%d]\n", redShirtSpeeds[i], blueShirtSpeeds[i])
		sum += max(redShirtSpeeds[i], blueShirtSpeeds[i])
	}

	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr1 := []int{3, 6, 7, 2, 1}
	arr2 := []int{5, 5, 3, 9, 2}
	fmt.Println(TandemBicycle(arr1, arr2, true))
	fmt.Println(TandemBicycle(arr1, arr2, false))

}
