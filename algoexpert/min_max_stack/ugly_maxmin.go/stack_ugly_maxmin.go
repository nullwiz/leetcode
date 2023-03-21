package ugly

import (
	"fmt"
)

type StackNode struct {
	Value int
	Next  *StackNode
}
type MinMaxStack struct {
	Top       *StackNode
	Min       int
	Max       int
	ArrayMins []int
	ArrayMaxs []int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{nil, 0, 0, []int{}, []int{}}

}

func (stack *MinMaxStack) Peek() int {
	fmt.Printf("top: %d\n", stack.Top.Value)
	return stack.Top.Value
}

func (stack *MinMaxStack) Pop() int {
	fmt.Println("popping val")
	if stack.Top != nil {
		if stack.Top.Value == stack.ArrayMaxs[len(stack.ArrayMaxs)-1] {
			if len(stack.ArrayMaxs) < 2 {
				stack.Max = 0
				stack.ArrayMaxs = []int{}
			} else {
				stack.ArrayMaxs = stack.ArrayMaxs[:len(stack.ArrayMaxs)-1]
				stack.Max = stack.ArrayMaxs[len(stack.ArrayMaxs)-1]
			}
		}
		if stack.Top.Value == stack.ArrayMins[len(stack.ArrayMins)-1] {
			if len(stack.ArrayMins) < 2 {
				stack.Min = 0
				stack.ArrayMins = []int{}
			} else {
				stack.ArrayMins = stack.ArrayMins[:len(stack.ArrayMins)-1]
				stack.Min = stack.ArrayMins[len(stack.ArrayMins)-1]
			}
		}

		ret := stack.Top.Value
		stack.Top = stack.Top.Next
		return ret
	} else {
		return -1
	}
}

func (stack *MinMaxStack) Push(number int) {
	// Write your code here.
	node := &StackNode{Value: number, Next: stack.Top}
	if stack.Top == nil {
		// Both values are min and max
		stack.ArrayMins = append(stack.ArrayMins, number)
		stack.Min = number
		stack.ArrayMaxs = append(stack.ArrayMaxs, number)
		stack.Max = number
	} else {
		if number < stack.Min {
			stack.ArrayMins = append(stack.ArrayMins, number)
			stack.Min = number
		}
		if number > stack.Max {
			stack.ArrayMaxs = append(stack.ArrayMaxs, number)
			stack.Max = number
		}
	}
	stack.Top = node
}

func (stack *MinMaxStack) GetMin() int {
	fmt.Printf("min:%d \n", stack.Min)
	return stack.Min
}

func (stack *MinMaxStack) GetMax() int {
	fmt.Printf("max:%d \n", stack.Max)
	return stack.Max
}

func (stack *MinMaxStack) Print() {
	fmt.Printf("stack : [")
	current := stack.Top
	for current != nil {
		fmt.Printf("%d  ", current.Value)
		current = current.Next
	}
	fmt.Printf("]\n")
	fmt.Printf("max: %d, min: %d\n", stack.Max, stack.Min)
}

func main() {
	stack := NewMinMaxStack()
	stack.Push(5)
	stack.GetMin()
	stack.GetMax()
	stack.Peek()
	stack.Push(7)
	stack.Push(2)
	stack.Push(-1)
	stack.Push(1)
	stack.GetMin()
	stack.GetMax()
	stack.Peek()
	stack.Push(2)
	stack.GetMin()
	stack.GetMax()
	stack.Peek()
	stack.Pop()
	stack.Pop()

	stack.GetMin()
	stack.GetMax()
	stack.Peek()
	stack.Pop()
	stack.GetMin()
	stack.GetMax()
	stack.Peek()

}
