package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func (tree *BinaryTree) insert(value int) {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BinaryTree{Value: value}
		} else {
			tree.Left.insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BinaryTree{Value: value}
		} else {
			tree.Right.insert(value)
		}
	}
}
func createBST(array []int) *BinaryTree {
	root := &BinaryTree{Value: array[0]}
	for i := 1; i < len(array); i++ {
		root.insert(array[i])
	}
	return root
}

func NodeDepths(root *BinaryTree, depth int) int {
	// Write your code here.
	if root == nil {
		return 0
	}
	left := NodeDepths(root.Left, depth+1)
	right := NodeDepths(root.Right, depth+1)
	return left + right
}
func minChange(n int, coins []int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}
	if len(coins) == 0 {
		return 0
	}
	count := 0
	// This will fail because it counts all non distinct combinations
	for i := 0; i < len(coins); i++ {
		if n >= coins[i] {
			count += minChange(n-coins[i], coins[i:])
		}
	}
	return count
}
func min_Change(n int, coins []int, memo []int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}
	if len(coins) == 0 {
		return 0
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = min_Change(n-coins[0], coins, memo) + min_Change(n, coins[1:], memo)
	return memo[n]

}

func CanDoSum(array []int, n int, i int) bool {
	if n == 0 {
		return true
	}
	if n < 0 || i < 0 {
		return false
	}
	result := CanDoSum(array, n, i-1) || CanDoSum(array, n-array[i], i-1)
	return result

}

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	return KnapsackHelper(items, capacity, 0)
}

func KnapsackHelper(items [][]int, capacity int, index int) []interface{} {
	if capacity == 0 {
		return []interface{}{0, items}
	}
	if capacity < 0 {
		return []interface{}{-9999, []int{0}}
	}
	if index == len(items)-1 {
		return []interface{}{capacity, items}
	}
	result_include := KnapsackHelper(items, capacity, index+1)
	result_exclude := KnapsackHelper(items, capacity-items[index][1], index+1)
	result_exclude[0] = result_exclude[0].(int) + items[index][0]
	if result_include[0].(int) > result_exclude[0].(int) {
		return result_include
	} else {
		return result_exclude
	}
}

func MaxSubsetSumNonAdjancent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return MaxSubsetSumNonAdjancentHelper(array, 0)
}

func MaxSubsetSumNonAdjancentHelper(array []int, i int) int {
	if i >= len(array) {
		return 0
	}
	// Include current number and skip the next
	sum1 := array[i] + MaxSubsetSumNonAdjancentHelper(array, i+2)
	sum2 := MaxSubsetSumNonAdjancentHelper(array, i+1)
	return max(sum1, sum2)
}

func countChangeWays(n int, coins []int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	ways := 0
	for i := 0; i < len(coins); i++ {
		ways += countChangeWays(n-coins[i], coins[i:])
	}
	return ways
}

func minCoinsForChange(n int, coins []int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return math.MaxInt32
	}
	c_min := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		min := 1 + minCoinsForChange(n-coins[i], coins)
		if min != math.MaxInt32 {
			if c_min > min {
				c_min = min
			}
		} else {
			return 0
		}
	}
	return c_min
}

func fib(n int) int {
	if n < 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

	fmt.Println(minCoinsForChange(10, []int{1, 3, 4}))
}
