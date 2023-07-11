package main

import "fmt"

var m = make(map[int]int)

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	if value, ok := m[n]; ok {
		return value
	}
	m[n] = climbStairs(n-1) + climbStairs(n-2)
	return climbStairs(n-1) + climbStairs(n-2)
}
func climbStairs2(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

func main() {
	fmt.Println(climbStairs(150))
	fmt.Println(climbStairs2(150))
}
