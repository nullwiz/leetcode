package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	// Create a frequency map.
	// O(n +m) | O(c)
	charmap := make(map[string]int)
	docmap := make(map[string]int)
	// O(n)
	for i := 0; i < len(characters); i++ {
		charmap[string(characters[i])]++
	}
	for i := 0; i < len(document); i++ {
		docmap[string(document[i])]++
	}

	for k := range docmap {
		if _, ok := charmap[k]; ok {
			if charmap[k] < docmap[k] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func GenDocSmaller(characters string, document string) bool {
	m := make(map[rune]int)
	for _, v := range characters {
		m[v]++
	}
	for _, v := range document {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
func main() {
	a := GenerateDocument("Bste!hetsi ogEAxpelrt x ", "AlgoExpert is the Best!")
	fmt.Println(a)
}
