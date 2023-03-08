package main

import (
	"testing"
	"time"
)

func TestSmallNums(t *testing.T) {
	arr := []string{"a", "b", "c", "e", "z", "k", "j"}
	// sorted should be
	expected := []string{"a", "b", "c", "e", "j", "k", "z"}
	// take time
	elapsed := time.Now()
	quicksort(arr)
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
