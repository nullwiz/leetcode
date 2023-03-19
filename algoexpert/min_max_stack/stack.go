package main

type minMax map[string][int]

type MinMaxStack struct {
	Size int
	Stack []int
	MinMax []minMax
}

func NewMinMaxStack() *NewMinMaxStack {
	return &NewMinMaxStack{
		Stack : []int{},
		MinMax: []minMax{},
	}
}

func (stack *MinMaxStack) Peek() int {
	return stack.Stack[stack.Size-1]
}

func (stack *MinMaxStack) Pop() int {
	ret := stack.Peek()
	stack.Size--
	stack.Stack = stack.Stack[:Stack.Size]
	stack.MinMax = stack.MinMax[:stack.Size]
	return val
}

func (stack *MinMaxStack){}
