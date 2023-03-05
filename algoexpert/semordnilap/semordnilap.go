package main

import "fmt"

func semordnilap(str []string) [][]string {
	m := make(map[string][]string)
	ret := [][]string{}
	for _, v := range str {
		if _, ok := m[reversed(v)]; ok {
			m[v] = []string{v, reversed(v)}
		} else {
			m[v] = append(m[v], v)
		}
	}
	for _, v := range m {
		if len(v) == 2 {
			ret = append(ret, v)
		}
	}
	return ret
}
func reversed(str string) string {
	b := len(str) - 1
	ret := make([]rune, len(str))
	for _, v := range str {
		ret[b] = v
		b--
	}
	return string(ret[:])
}
func main() {
	words := []string{"dog", "desserts", "god", "stressed"}
	fmt.Println(semordnilap(words))
}
