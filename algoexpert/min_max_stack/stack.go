package main

import "fmt"

type minMax map[string]int

type MinMaxStack struct {
	Size   int
	Stack  []int
	MinMax []minMax
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		Stack:  []int{},
		MinMax: []minMax{},
	}
}

func (stack *MinMaxStack) Peek() int {
	return stack.Stack[stack.Size-1]
}

func (stack *MinMaxStack) Pop() int {
	ret := stack.Peek()
	stack.Size--
	stack.Stack = stack.Stack[:stack.Size]
	stack.MinMax = stack.MinMax[:stack.Size]
	return ret
}

func (stack *MinMaxStack) Push(number int) {
	currMinMax := minMax{}
	stack.Size++
	if stack.Size == 1 {
		currMinMax["min"], currMinMax["max"] = number, number
	} else {
		currMinMax["min"] = min(stack.MinMax[stack.Size-2]["min"], number)
		currMinMax["max"] = max(stack.MinMax[stack.Size-2]["max"], number)
	}
	stack.Stack = append(stack.Stack, number)
	stack.MinMax = append(stack.MinMax, currMinMax)
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func (stack *MinMaxStack) String() string {
	return fmt.Sprintf("MinMaxStack {\n Size: %d \n Stack [ %v ]\n Min: %d\tMax:%d\n}", stack.Size, stack.Stack, stack.MinMax[stack.Size-1]["min"], stack.MinMax[stack.Size-1]["max"])
}

func main() {
	stack := NewMinMaxStack()
	stack.Push(3)
	fmt.Println(stack)
	stack.Push(1)
	stack.Push(-1)
	stack.Push(51)
	stack.Push(11)
	stack.Push(999)
	stack.Push(2)
	fmt.Println(stack)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack)
}
