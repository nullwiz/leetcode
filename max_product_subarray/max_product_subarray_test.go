package main

import (
	"testing"
)

// [2,3,4,5,6,-2,-3]  [1,2,-2,4,5,1-3] [-1,-2,3,4,5,6] [-2,-2,-3,-4,5,6] flip_c = even
// [-2,-4,-3,2,5] [1,2,-4,-2,5,-3] flic =  odd
// [2,3,4,5,5]
// [0, 0, 1, 1 ,0, 1, 1, 1, 0, 0, 1, 1]
///

func TestMaxProductEvenNegatives(t *testing.T) {
	{
		a := []int{-2, 3, -4, 2, 3, 4, 5, 6, -2, -4}
		expected := 138240
		result := maxProduct(a)
		if result != expected {
			t.Errorf("We failed!! result was %d, we expected %d", result, expected)
		}
	}
}

func TestMaxProductOddNegatives(t *testing.T) {
	{
		a := []int{-2, 3, -1, -3}
		expected := 9
		result := maxProduct(a)
		if result != expected {
			t.Errorf("We failed!! result was %d, we expected %d", result, expected)
		}
	}
}
