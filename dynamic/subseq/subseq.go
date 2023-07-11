package main

import (
	"fmt"
	"math"
)

func isSubsequence(s string, t string) bool {
	lt := len(t)
	ls := len(s)
	i, j := 0, 0
	for i < lt && j < ls {
		if s[j] == t[i] {
			j++
		}
		i++
	}
	return j == ls
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MaxSubsetSumNonAdjancentHelper(array []int, i int) int {
	i_sum, e_sum := 0, 0
	if len(array) == 0 {
		return 0
	}
	if i > len(array)-1 {
		return 0
	}
	i_sum += MaxSubsetSumNonAdjancentHelper(array, i+2) + array[i]
	e_sum += MaxSubsetSumNonAdjancentHelper(array, i+1)
	return max(i_sum, e_sum)
}
func MaxSubsetSumNonAdjancent(array []int) int {
	return MaxSubsetSumNonAdjancentHelper(array, 0)
}
func MaxSubsetSumNonAdjancentHelperIndexes(array []int, i int, results []int) (int, []int) {
	i_sum, e_sum := 0, 0

	if i > len(array)-1 {
		return 0, results
	}
	currInclusionResults, currExclusionResults := []int{}, []int{}
	currInclusionResults = append(currInclusionResults, array[i])

	inclusionValue, inclusionResults := MaxSubsetSumNonAdjancentHelperIndexes(array, i+2, []int{})
	i_sum += inclusionValue + array[i]
	currInclusionResults = append(currInclusionResults, inclusionResults...)
	exclusionValue, exclusionResults := MaxSubsetSumNonAdjancentHelperIndexes(array, i+1, []int{})
	e_sum += exclusionValue
	currExclusionResults = append(currExclusionResults, exclusionResults...)

	if i_sum > e_sum {
		return i_sum, currInclusionResults
	}
	return e_sum, currExclusionResults
}
func MaxSubsetSumNonAdjancentIndexes(array []int) (int, []int) {
	return MaxSubsetSumNonAdjancentHelperIndexes(array, 0, []int{})
}

func EasyMaxSubsetSumIndexes(array []int) (int, []int) {
	excl := 0
	incl := array[0]
	exc_indices := []int{}
	inc_indices := []int{array[0]}
	for i := 1; i < len(array); i++ {
		new_excl := max(excl, incl)
		new_excl_indices := []int{}
		if new_excl == excl {
			new_excl_indices = exc_indices
		} else {
			new_excl_indices = inc_indices
		}
		incl = excl + array[i]
		inc_indices = append(exc_indices, array[i])
		excl = new_excl
		exc_indices = append(exc_indices, new_excl_indices...)
	}
	if excl > incl {
		return excl, exc_indices
	}
	return incl, inc_indices
}

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	m := make(map[int]int)
	val := MinNumberOfCoinsForChangeHelper(n, denoms, m)
	return val
}

func MinNumberOfCoinsForChangeHelper(n int, denoms []int, memo map[int]int) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return 1234
	}
	if val, ok := memo[n]; ok {
		return val
	}
	currMin := 1234

	for i := 0; i < len(denoms); i++ {
		if n >= denoms[i] {
			recursionMin := MinNumberOfCoinsForChangeHelper(n-denoms[i], denoms, memo)
			if recursionMin != -1 && recursionMin < 1234 {
				currMin = recursionMin + 1
			}
		}
	}
	if currMin != 1234 {
		memo[n] = currMin
	} else {
		return -1
	}
	return currMin
}

func MaxSumIncreasingSubseq(array []int) (int, []int) {

	return -1, []int{}
}

func findLargestSubseq(arr []int) ([]int, int) {
	le := make([]int, len(arr))
	copy(le, arr)
	indices := make([][]int, len(arr))
	for i := 1; i < len(arr); i++ {
		for k := 0; k < i; k++ {
			if arr[k] < arr[i] {
				le[i] = max(le[i], le[k]+arr[i])
			}
		}
	}
	maxim := 0
	index := 0
	for k, v := range le {
		maxim = max(maxim, v)
		index = k
	}
	return indices[index], maxim

}

func findLargestSumSubseq(arr []int) (int, []int) {
	memo := make([][]*int, len(arr))
	for i := range memo {
		memo[i] = make([]*int, len(arr))
	}
	maxNum := math.MinInt32 // keep track of the maximum number
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}
	prev := -1
	maxSum, _ := findLargestSumSubseqHelper(arr, 0, prev, maxNum, memo)
	return maxSum, nil
}

func findLargestSumSubseqHelper(arr []int, i int, prev int, maxNum int, memo [][]*int) (int, []int) {
	// Base case: if we've traversed the entire array
	if i == len(arr) {
		if prev != -1 {
			return 0, nil
		}
		return maxNum, nil // return the maximum number if no number was added
	}
	if prev != -1 && memo[i][prev] != nil {
		return *memo[i][prev], nil
	}
	// Exclude current number: use the existing results slice
	exclude, _ := findLargestSumSubseqHelper(arr, i+1, prev, maxNum, memo)

	// If the current number can be included (greater than previous), create a new results slice and add it
	include := exclude
	if prev == -1 || arr[i] > arr[prev] {
		newInclude, _ := findLargestSumSubseqHelper(arr, i+1, i, maxNum, memo)
		include = max(arr[i]+newInclude, include)
	}
	result := max(include, exclude)
	if prev != -1 {
		memo[i][prev] = &result
	}
	return result, nil
}
func sum(arr []int) int {
	count := 0
	for _, v := range arr {
		v += count
	}
	return count
}

func main() {
	//fmt.Println(MaxSumIncreasingSubseq([]int{10, 30, 5, 2, 50, 60}))
	arr := []int{-5, -4, -3, -2, -1}
	//arr := []int{10, 70, 20, 50}
	//arr := []int{1, 2, 3}
	fmt.Println(findLargestSumSubseq(arr))
}
