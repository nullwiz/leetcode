package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// handle edge case, only two values
	x := 0
	newmap := make(map[int]int)
	for k, v := range nums {
		// x = z - y
		x = target - v
		// if x exists in hashmap, return x and y
		if _, ok := newmap[x]; ok {
			return []int{newmap[x], k}
		}
		// if x does not exist
		newmap[v] = k
	}
	return nil
}

func main() {
	tc := []int{3, 2, 3}
	a := twoSum(tc, 6)
	fmt.Println(a)
}
