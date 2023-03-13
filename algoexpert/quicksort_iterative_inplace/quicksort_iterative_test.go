package main

import (
	"testing"
	"time"
)

func TestSmallNums(t *testing.T) {
	arr := []int{1, 5, 6, 7, 8, 2, 56, -2, 3, 1}
	// sorted should be
	expected := []int{-2, 1, 1, 2, 3, 5, 6, 7, 8, 56}
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
}

func TestBigNums1k(t *testing.T) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}

func TestBigNums2k(t *testing.T) {
	arr := make([]int, 2000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}

func TestBigNums3k(t *testing.T) {
	arr := make([]int, 4000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}
func TestBigNums4k(t *testing.T) {
	arr := make([]int, 8000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}
func TestBigNums5k(t *testing.T) {
	arr := make([]int, 16000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}

func TestBigNums10k(t *testing.T) {
	arr := make([]int, 32000)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}


func TestBigNums50k(t *testing.T) {
	arr := make([]int, 32000 * 5)
	for i := range arr {
		arr[i] = 1
	}
	// sorted should be
	expected := arr
	// take time
	elapsed := time.Now()
	quicksortIterative(arr)
	// compare slices
	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	}
	//compare time
	if time.Since(elapsed) > time.Millisecond*100 {
		t.Errorf("Expected less than 100ms, got %v", time.Since(elapsed))
	}
	t.Log("Elapsed time: ", time.Since(elapsed))
}