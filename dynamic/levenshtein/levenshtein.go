package main

import (
	"fmt"
	"strconv"
)

func LevenshteinDistance(a, b string) int {
	if a == b {
		return 0
	}
	rows := len(a) + 1
	columns := len(b) + 1
	ret := make([][]int, rows)
	for i := range ret {
		ret[i] = make([]int, columns)
	}
	// Populate last column
	for j := 0; j < rows; j++ {
		ret[j][columns-1] = rows - j - 1
	}
	// populate last row
	for j := 0; j < columns; j++ {
		ret[rows-1][j] = columns - j - 1
	}
	// Populate each column, from right to left
	for k := rows - 2; k >= 0; k-- {
		for l := columns - 2; l >= 0; l-- {
			if a[k] == b[l] {
				ret[k][l] = ret[k+1][l+1]
			} else {
				val := 1 + min(ret[k][l+1], ret[k+1][l], ret[k+1][l+1])
				ret[k][l] = val
			}
		}
	}

	return ret[0][0]
}

// Starts at the end of the string.
func LevenshteinDistHelper(a string, b string, i int, j int, memo map[string]int) int {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	key := strconv.Itoa(i) + "," + strconv.Itoa(j)
	if value, ok := memo[key]; ok {
		return value
	}
	if i == 0 {
		return 1
	}
	if j == 0 {
		return 1
	}
	if a[i-1] == b[j-1] {
		return LevenshteinDistHelper(a, b, i-1, j-1, memo)
	}
	// They are not the same.
	return 1 + min(LevenshteinDistHelper(a, b, i, j-1, memo),
		LevenshteinDistHelper(a, b, i-1, j, memo),
		LevenshteinDistHelper(a, b, i-1, j-1, memo))
}

func LevenshteinDist(a string, b string) int {
	memo := make(map[string]int)
	return LevenshteinDistHelper(a, b, len(a), len(b), memo)
}

func min(x int, y int, z int) int {
	if x < y && x < z {
		return x
	} else if y < x && y < z {
		return y
	} else {
		return z
	}
}

// returns map of missing characters from a to make b
func missingChars(a string, b string) map[string]int {
	freqa := make(map[string]int)
	freqb := make(map[string]int)
	ret := make(map[string]int)

	for _, char := range a {
		freqa[string(char)]++
	}

	for _, char := range b {
		freqb[string(char)]++
	}

	for char, count := range freqb {
		if val, ok := freqa[char]; ok {
			diff := count - val
			if diff > 0 {
				ret[char] = diff
			}
		} else {
			diff := count
			ret[char] = diff
		}
	}
	return ret
}

func pprint(a [][]int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	str1 := "abbb"
	str2 := "bbba"
	fmt.Println(LevenshteinDist(str1, str2))
	str1 = "abc"
	str2 = "yabd"
	fmt.Println(LevenshteinDist(str1, str2))
}
