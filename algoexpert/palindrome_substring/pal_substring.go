package main

import (
	"fmt"
	"strings"
)

func LongestPalindromicSubstring(str string) string {
	//	a, b := 0, len(str)-1
	c := []string{}
	soFar := ""
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			c = append(c, str[i:j])
		}
	}
	fmt.Println(c)
	for _, v := range c {
		if reversed(v) == true && len(v) > len(soFar) {
			soFar = v
		}
	}
	fmt.Println(soFar)
	return soFar
}

func reversed(str string) bool {
	r := make([]string, len(str))
	a := len(str) - 1
	for i := 0; i < len(str); i++ {
		r[i] = string(str[a])
		a--
	}
	b := strings.Join(r, "")
	fmt.Printf("String looks like :%s\n", b)
	return str == b
}

func main() {
	a := "neuquenasd1234"
	LongestPalindromicSubstring(a)
}
