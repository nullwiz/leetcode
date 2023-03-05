package main

import "testing"

func TestSquareTraverse(t *testing.T) {
	arr := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	actual := SpiralTraverse(arr)
	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestNonSquareTraverse(t *testing.T) {
	arr := [][]int{{1, 2, 3, 4}, {10, 11, 12, 5}, {9, 8, 7, 6}}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	actual := SpiralTraverse(arr)
	if len(expected) != len(actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}

}
