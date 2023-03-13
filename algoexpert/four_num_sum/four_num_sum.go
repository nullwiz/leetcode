package main

import "fmt"

func FourNumberSum(array []int, target int) [][]int {
	m := make(map[int][][]int)
	ret := [][]int{}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			diff := target - currentSum
			if _, ok := m[diff]; ok {
				for _, v := range m[diff] {
					pair := []int{array[i], array[j]}
					pair = append(pair, v...)
					ret = append(ret, pair)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			if _, ok := m[currentSum]; ok {
				m[currentSum] = append(m[currentSum], []int{array[i], array[k]})
			} else {
				m[currentSum] = append(m[currentSum], []int{array[i], array[k]})
			}
		}

	}
	return ret
}

func main() {
	arr := []int{7, 6, 4, -1, 1, 2}
	fmt.Println(FourNumberSum(arr, 16))
}
