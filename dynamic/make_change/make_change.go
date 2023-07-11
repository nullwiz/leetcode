package main

import (
	"fmt"
)

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	if n == 0 || len(denoms) == 0 {
		return -1
	}
	ways := make([]int, n+1)
	ways[0] = 1
	for i := 0; i < len(denoms); i++ {
		for j := 1; j < len(ways); j++ {
			if denoms[i] <= j {
				ways[j] += ways[j-denoms[i]]
			}
		}
	}
	fmt.Println(ways)
	return ways[n]

}

var memoizer = make(map[int]int)

func NumberOfWaysToMakeMinChange(n int, denoms []int) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return 333333
	}
	if val, ok := memoizer[n]; ok {
		return val
	}

	minCoins := 333333
	for _, denom := range denoms {
		if denom <= n {
			res := NumberOfWaysToMakeMinChange(n-denom, denoms)
			if res != -1 && res < minCoins {
				memoizer[n] = res + 1
				minCoins = res + 1
			}
		}
	}

	if minCoins == 333333 {
		return -1
	}
	return memoizer[n]
}

func minList(list []int) int {
	min := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		}
	}
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// This one was needed because the workspace was
// having trouble with the memoized map out of the main function
func _MinNumberOfCoinsForChange(n int, denoms []int) int {
	memoizer := make(map[int]int)

	var helperFunc func(int) int
	helperFunc = func(n int) int {
		if n == 0 {
			return 0
		} else if n < 0 {
			return 333333
		}
		if val, ok := memoizer[n]; ok {
			return val
		}

		minCoins := 333333
		for _, denom := range denoms {
			if denom <= n {
				res := helperFunc(n - denom)
				if res != -1 && res < minCoins {
					minCoins = res + 1
				}
			}
		}
		// update memoizer map
		if minCoins != 333333 {
			memoizer[n] = minCoins
		} else {
			return -1
		}
		return minCoins
	}

	return helperFunc(n)
}
func NumberOfWaysToTraverseGraph(width int, height int) int {
	graph := make([][]int, height)
	for i := range graph {
		graph[i] = make([]int, width)
	}
	val := traverseGraph(graph, 0, 0)
	return val
}

func traverseGraph(graph [][]int, x int, y int) int {
	if x == len(graph)-1 && y == len(graph[0])-1 {
		// we are at bottom right
		return 1
	}
	if x >= len(graph) || x < 0 || y >= len(graph[0]) || y < 0 {
		return 0
	}
	if graph[x][y] != 0 {
		return graph[x][y]
	}
	graph[x][y] = traverseGraph(graph, x+1, y) + traverseGraph(graph, x, y+1)

	return graph[x][y]
}
func MinNumberOfWaysToTraverseGraph(width int, height int) int {
	graph := make([][]int, height)
	for i := range graph {
		graph[i] = make([]int, width)
	}
	val := traverseGraphMin(graph, 0, 0)
	return val
}

func traverseGraphMin(graph [][]int, x int, y int) int {
	if x == len(graph)-1 && y == len(graph[0])-1 {
		// we are at bottom right
		return 1
	}
	if x >= len(graph) || x < 0 || y >= len(graph[0]) || y < 0 {
		return 0
	}
	if graph[x][y] != 0 {
		return graph[x][y]
	}
	graph[x][y] = 1 + min(traverseGraph(graph, x+1, y), traverseGraph(graph, x, y+1))
	return graph[x][y]
}

func NumberOfWaysToTraverseGraph1(width int, height int) int {
	// Write your code here.
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)

	}
	return helper(grid, 0, 0)
}

func helper(grid [][]int, x int, y int) int {
	if x > len(grid)-1 || y > len(grid[0])-1 {
		return 0
	}
	if x == len(grid)-1 && y == len(grid[0])-1 {
		return 1
	}
	grid[x][y] = helper(grid, x+1, y) + helper(grid, x, y+1)
	return grid[x][y]
}

func NumberOfWaysToMakeChange1(n int, denoms []int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	ways := 0
	for i := 0; i < len(denoms); i++ {
		if denoms[i] <= n {
			ways += NumberOfWaysToMakeChange1(n-denoms[i], denoms[i:])
		}
	}
	return ways
}

func MinNumberOfWaysToMakeChange(n int, denoms []int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if len(denoms) == 0 {
		return 0
	}
	ways := MinNumberOfWaysToMakeChange(n-denoms[0], denoms)
	ways2 := MinNumberOfWaysToMakeChange(n, denoms[1:])

	return ways + ways2
}

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(NumberOfWaysToTraverseGraph1(4, 3))
	fmt.Println(MinNumberOfWaysToMakeChange(3, []int{1, 2, 3}))
}
