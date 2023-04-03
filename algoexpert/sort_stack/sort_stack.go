package main

import "fmt"

func SortStack(stack []int) []int {
	// Write your code here.
	if len(stack) <= 1 {
		return stack
	}
	firstValue := Pop(&stack)
	SortStack(stack)
	secondValue := Pop(&stack)

	var largerValue int
	var smallerValue int

	if firstValue > secondValue {
		largerValue = firstValue
		smallerValue = secondValue
	} else {
		largerValue = secondValue
		smallerValue = firstValue
	}
	Push(&stack, largerValue)
	SortStack(stack)
	Push(&stack, smallerValue)
	return stack
}

func Pop(stack *[]int) int {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func Peek(stack []int) int {
	return stack[len(stack)-1]
}

func Push(stack *[]int, item int) {
	*stack = append(*stack, item)
}

func main() {
	stack := []int{5, 2, -2, 4, 3, 1}
	fmt.Println(SortStack(stack))

}
