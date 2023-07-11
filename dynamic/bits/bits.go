package main

import "fmt"

func countBits2(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	for i := 1; i <= n; i++ {
		ans[i] = bitCount(i)
	}
	return ans
}

// dynamic solution
func countBits(n int) []int {
	dp := make([]int, n+1)
	offset := 1

	for i := 1; i < n+1; i++ {
		if offset*2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i-offset]
	}

	return dp
}

func numToBinary(n int) int {
	count := 0
	for n != 0 {
		v := n % 2
		if v > 0 {
			count++
		}
		n = n / 2
	}
	return count
}

func bitCount(n int) int {
	cnt := 0
	for n != 0 {
		n &= n - 1
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(countBits(5))
	fmt.Println(countBits2(5))
}
