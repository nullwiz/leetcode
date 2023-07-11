package main

import "fmt"

func pascalTriangle(row int) [][]int {
	ret := [][]int{}
	if row < 0 {
		return ret
	}
	for i := 0; i <= row; i++ {
		if i == 0 {
			continue
		}
		if i == 1 {
			ret = append(ret, []int{1})
			continue
		}
		if i == 2 {
			ret = append(ret, []int{1, 1})
			continue
		}
		last_row := ret[len(ret)-1]
		new_row := make([]int, len(last_row)+1)
		new_row[0] = 1
		new_row[len(new_row)-1] = 1
		for j := 0; j < len(last_row)-1; j++ {
			new_row[j+1] = last_row[j] + last_row[j+1]
		}
		ret = append(ret, new_row)
	}
	return ret
}

func pascalTriangeIndexed(row int) []int {
	ret := make([]int, row+1)
	ret[0] = 1
	for i := 1; i <= row; i++ {
		for j := i; j > 0; j-- {
			if j == i {
				ret[j] = 1
			} else {
				ret[j] = ret[j] + ret[j-1]
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(pascalTriangeIndexed(5))
}
