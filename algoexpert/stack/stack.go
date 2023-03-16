package main

import (
	"errors"
	"fmt"
)

// Stack Interface
type StackInterface interface {
	pop() (int, error)
	push(item int) error
	isEmpty() bool
	peek() int
	Fill(arr []int)
	Print()
}

// Stack implementation using a linked list
type StackNode struct {
	Value int
	Next  *StackNode
}
type Stack struct {
	Top *StackNode
}

func (f *Stack) pop() (int, error) {
	if f.Top == nil {
		return 0, errors.New("empty stack")
	}
	item := f.Top.Value
	f.Top = f.Top.Next
	return item, nil
}

func (f *Stack) push(item int) error {
	node := &StackNode{Value: item}
	node.Next = f.Top
	f.Top = node
	return nil
}

func (f *Stack) isEmpty() bool {
	return f.Top == nil
}

func (f *Stack) peek() int {
	return f.Top.Value
}

func (f *Stack) Print() {
	fmt.Printf("Stack: ")
	for current := f.Top; current != nil; current = current.Next {
		fmt.Printf("%d ", current.Value)
	}
	fmt.Println()
}
func (f *Stack) Fill(arr []int) {
	for _, v := range arr {
		f.push(v)
	}
}
func (f *Stack) String() string {
	return fmt.Sprintf("Stack{ Type: %T , Top: %p }\n", f, &f.Top)
}

// Stack implementation using an array
type ArrayStack struct {
	arr []int
}

func (s *ArrayStack) pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("Stack is empty")
	}
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item, nil
}

func (s *ArrayStack) push(item int) error {
	s.arr = append(s.arr, item)
	return nil
}

func (s *ArrayStack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *ArrayStack) peek() int {
	return s.arr[len(s.arr)-1]
}

func (s *ArrayStack) Fill(arr []int) {
	for _, v := range arr {
		s.push(v)
	}
}

func (s *ArrayStack) Print() {
	fmt.Printf("ArrayStack: ")
	for i := len(s.arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", s.arr[i])
	}
	fmt.Println()
}

func (s *ArrayStack) String() string {
	return fmt.Sprintf("ArrayStack{ Type: %T , arr: %p }\n", s, &s.arr)
}

func main() {
	var stackInterface StackInterface
	// Create a stack using a linked list
	var stack Stack
	// Assign the stack to the interface
	stackInterface = &stack
	stackInterface.push(1)
	arr := []int{1, 2, 3, 4, 5}
	stackInterface.Fill(arr)
	fmt.Print(&stack)
	stackInterface.Print()
	// Peek the top
	fmt.Printf("top %d\n", stackInterface.peek())
	// Pop the top
	fmt.Println(stackInterface.pop())

	// Create a stack using an array
	var arrayStack ArrayStack
	// Assign the stack to the interface
	stackInterface = &arrayStack
	stackInterface.push(1)
	fmt.Print(&arrayStack)
	stackInterface.Fill(arr)
	// Peek the top
	fmt.Printf("top %d\n", stackInterface.peek())
	// This implementation prints differntly
	arrayStack.Print()
	fmt.Println(stackInterface.pop())

}
