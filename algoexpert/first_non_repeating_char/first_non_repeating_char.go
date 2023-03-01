package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	// Map of runes with an array of two ints
	m := make(map[rune][2]int)
	for k, v := range str {
		// This saves the occurence.
		m[v] = [2]int{m[v][0] + 1, k}
	}
	ret := len(str)
	for _, v := range m {
		if v[0] == 1 {
			if v[1] < ret {
				ret = v[1]
			}
		}
	}
	if ret == len(str) {
		return -1
	}
	return ret
}

func main() {
	str := "🔨abcdcaf"
	fmt.Println(FirstNonRepeatingCharacter(str))
}
