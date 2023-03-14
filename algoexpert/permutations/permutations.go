package main

import "fmt"

func GetPermutations(array []int) [][]int {
	ret := [][]int{}
	getPermutations(array, []int{}, &ret)
	return ret
}

func getPermutations(array []int, perm []int, perms *[][]int) {
	if len(array) == 0 {
		*perms = append(*perms, perm)
	}
	for k, v := range array {
		// newArray := []int{}
		// newArray = append(newArray, array[:k]...)
		// newArray = append(newArray, array[k+1:]...)
		// newPerm := []int{}
		// newPerm = append(newPerm, perm...)
		// newPerm = append(newPerm, v)
		// getPermutations(newArray, newPerm, perms)
		getPermutations(append(append([]int{}, array[:k]...), array[k+2:]...), append(perm, v), perms)
	}
}

func main() {
	array := []int{}
	fmt.Println(GetPermutations(array))
}
