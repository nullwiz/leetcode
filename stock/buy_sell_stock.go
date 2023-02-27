package main

import (
	"fmt"
)

// [1,2,3,4]
// [0] = 2*3*4
// [1] = 1*3*4
// [2] = 1*2*4
// [3] = 1*2*3

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	a := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(a))
}

// [1,2,3,4,5,6]
// sort the array from lowest price to highest price
// find min index left
// find max index right os that minIn < maxIn
