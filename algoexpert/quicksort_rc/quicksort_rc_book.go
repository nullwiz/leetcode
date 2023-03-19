package main

import "math/rand"

func quicksort_book(v []int, left int, right int) {
	if left >= right {
		return
	}
	swap(v, left, rand.Int()%right)
	last := left
	for i := left + 1; i <= right; i++ {
		if v[i] < v[left] {
			last++
			swap(v, last, i)
		}
	}
	swap(v, left, last)
	quicksort_book(v, left, last-1)
	quicksort_book(v, last+1, right)
}

func swap(v []int, i int, j int) {
	v[i], v[j] = v[j], v[i]
}
