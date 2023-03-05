package main

import "fmt"

func SpiralTraverse(array [][]int) []int {
	start_col := len(array[0]) - 1
	if len(array) != len(array[0]) {
		start_col -= 1
	}
	end_row := len(array) / 2
	end_col := len(array[0])/2 - 1
	b := PrintPerimeter(array, 0, start_col, end_row, end_col)
	return b
}

func PrintPerimeter(array [][]int, start_row int, start_col int, end_row int, end_col int) []int {
	ret := []int{}
	if start_row == end_row && start_col == end_col {
		ret = append(ret, array[start_row][start_col], array[start_row][start_col+1])
		return ret
	}
	// Traverse the top side
	for i := start_row; i < len(array[start_row])-start_row-1; i++ {
		ret = append(ret, array[start_row][i])
	}
	// Traverse the right side
	for i := start_row; i < len(array)-start_row-1; i++ {
		ret = append(ret, array[i][start_col])
	}
	// Traverse the bottom side
	for i := start_col; i > start_row; i-- {
		ret = append(ret, array[start_col][i])
	}
	// Traverse the left side
	for i := start_col; i > start_row; i-- {
		ret = append(ret, array[i][start_row])
	}
	fmt.Println(ret)
	ret = append(ret, PrintPerimeter(array, start_row+1, start_col-1, end_row, end_col)...)
	return ret
}
func main() {
	arr := [][]int{{1, 2, 3, 4}, {10, 11, 12, 5}, {9, 8, 7, 6}}
	fmt.Println(SpiralTraverse(arr))
}
