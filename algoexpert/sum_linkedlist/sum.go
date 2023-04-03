package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	arr1 := ReturnIntegerForList(linkedListOne)
	arr2 := ReturnIntegerForList(linkedListTwo)
	result := arr1 + arr2
	a := strconv.Itoa(result)
	res := strings.Split(a, "")
	res = ReverseArr(res, len(res)-1)
	fmt.Println(res)
	return CreateLinkedList(res)
}

func ReturnIntegerForList(ll *LinkedList) int {
	arr := []string{}
	for ll != nil {
		arr = append(arr, strconv.Itoa(ll.Value))
		ll = ll.Next
	}
	return ReverseArrAndReturnInteger(arr, len(arr)-1)
}

func ReverseArrAndReturnInteger(arr []string, length int) int {
	ret := ReverseArr(arr, len(arr)-1)
	ok, _ := strconv.Atoi(strings.Join(ret[:], ""))
	return ok
}

func ReverseArr(arr []string, length int) []string {
	ret := []string{}
	for i := length; i >= 0; i-- {
		ret = append(ret, arr[i])
	}
	return ret
}
func CreateLinkedList(arr []string) *LinkedList {
	head := &LinkedList{Value: 0, Next: nil}
	node := head
	for i, v := range arr {
		val, _ := strconv.Atoi(v)
		node.Value = val
		if i < len(arr)-1 {
			newNode := &LinkedList{Value: 0, Next: nil}
			node.Next = newNode
			node = node.Next
		}
	}
	return head
}
