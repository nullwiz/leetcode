package main

import "fmt"

// Buildings can only see the sunset if their height is strictly
// increasing from the direction specified
type SunsetStack struct {
	Stack         []int
	Size          int
	SunsetHeights []int
	LocalMax      int
}

func SunsetHelper(buildings []int, direction string) []int {
	if len(buildings) < 1 {
		return []int{}
	}
	stack := &SunsetStack{Stack: []int{}, Size: 0}
	stack.Fill(buildings, direction)
	if direction == "EAST" {
		reversedHeights := []int{}
		for i := len(stack.SunsetHeights) - 1; i >= 0; i-- {
			reversedHeights = append(reversedHeights, stack.SunsetHeights[i])
		}
		return reversedHeights
	}
	return stack.SunsetHeights
}

func (stack *SunsetStack) Push(number int) {
	stack.Stack = append(stack.Stack, number)
	stack.Size++
}

func (stack *SunsetStack) Fill(buildings []int, direction string) {
	// Fill the stack based on the direction and add local max values to the sunsetheight's list.
	// west = left; east = right.
	if direction == "EAST" {
		for i := len(buildings) - 1; i >= 0; i-- {
			stack.Push(buildings[i])
			if buildings[i] > stack.LocalMax {
				// This building can see the sunset.
				stack.SunsetHeights = append(stack.SunsetHeights, i)
			}
			// Update the current max.
			stack.LocalMax = max(stack.LocalMax, buildings[i])
		}
	}
	if direction == "WEST" {
		for i := 0; i < len(buildings); i++ {
			stack.Push(buildings[i])
			if buildings[i] > stack.LocalMax {
				// This building can see the sunset.
				stack.SunsetHeights = append(stack.SunsetHeights, i)
			}
			// Update the current max.
			stack.LocalMax = max(stack.LocalMax, buildings[i])
		}
	}
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func SunsetViews(buildings []int, direction string) []int {
	return SunsetHelper(buildings, direction)
}

func main() {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"
	fmt.Println(SunsetViews(buildings, direction))
}
