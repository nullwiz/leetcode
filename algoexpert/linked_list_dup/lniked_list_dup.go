package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	head := linkedList
	for head.Next != nil {
		if head.Next.Value == head.Value {
			if head.Next.Next != nil {
				head.Next = head.Next.Next
			} else {
				head.Next = nil
			}
		} else {
			head = head.Next
		}
	}
	return linkedList
}

func appendNodeAtTail(head *LinkedList, value int) *LinkedList {
	if head == nil {
		head = &LinkedList{value, nil}
		return head
	}
	node := head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &LinkedList{value, nil}
	return head
}

func main() {
	head := &LinkedList{Value: 1}
	// head.next = &LinkedList{value: 2}
	// head.next.next = &LinkedList{value: 3}
	// head.next.next.next = &LinkedList{value: 3}
	// head.next.next.next.next = &LinkedList{value: 3}
	// head.next.next.next.next.next = &LinkedList{value: 3}
	// head.next.next.next.next.next.next = &LinkedList{value: 3}
	// head.next.next.next.next.next.next.next = &LinkedList{value: 3}
	a := []int{1, 3, 4, 4, 4, 5, 6, 6}
	for _, v := range a {
		appendNodeAtTail(head, v)

	}
	RemoveDuplicatesFromLinkedList(head)
}
