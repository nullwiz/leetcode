package main

import "fmt"

type BracketStack struct {
	Stack    []string
	Size     int
	Closings int
}

func NewBracketStack() *BracketStack {
	return &BracketStack{Stack: []string{}, Size: 0, Closings: 0}
}
func (stack *BracketStack) Push(s string) {
	stack.Size++
	stack.Stack = append(stack.Stack, s)
}
func (stack *BracketStack) isBalanced() bool {
	return stack.Size == 0
}

func (stack *BracketStack) Peek() string {
	return stack.Stack[stack.Size-1]
}

func (stack *BracketStack) Pop() string {
	if stack.Size == 0 {
		return "0"
	}
	val := stack.Peek()
	stack.Size--
	stack.Stack = stack.Stack[:stack.Size]
	fmt.Println("Stack is:")
	fmt.Println(stack.Stack)
	return val
}

func isClosing(v rune) bool {
	if v == '}' || v == ']' || v == ')' {
		return true
	}
	return false
}
func isOpening(v rune) bool {
	if v == '{' || v == '[' || v == '(' {
		return true
	}
	return false
}

func BalancedBrackets(s string) bool {
	newStack := NewBracketStack()
	for _, v := range s {
		fmt.Println(string(v))
		if isClosing(v) {
			if newStack.Size == 0 {
				return false
			}
			open := newStack.Pop()
			if v == '}' && open != "{" || v == ']' && open != "[" || v == ')' && open != "(" {
				return false
			}
		}
		if isOpening(v) {
			newStack.Push(string(v))
		}
	}
	if newStack.isBalanced() {
		return true
	}
	return false
}

func main() {
	s := "{[[[[({(}))]]]]}"
	fmt.Println(BalancedBrackets(s))
}
