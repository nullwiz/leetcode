// Feel free to add methods and fields to the struct definitions.
package linkedList

import "fmt"

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func bindNodes(nodeOne *Node, nodeTwo *Node) {
	nodeOne.Next = nodeTwo
	nodeTwo.Prev = nodeOne
}
func NewNode(value int) *Node { return &Node{Value: value} }

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil}
}

func (node *Node) Disconnect() {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil

}
func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	nodeToInsert.Disconnect()
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node
	if node.Prev != nil {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}
	nodeToInsert.Disconnect()
	nodeToInsert.Next = node.Next
	nodeToInsert.Prev = node
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}

	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	node := ll.Head
	for i := 1; i <= position; i++ {
		node = node.Next
	}
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	for node := ll.Head; node != nil; node = node.Next {
		if node.Value == value {
			node.Disconnect()
			if node == ll.Head {
				ll.Head = ll.Head.Next
			}
			if node == ll.Tail {
				ll.Tail = ll.Tail.Prev
			}
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	node.Disconnect()
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	for node := ll.Head; node != nil; node = node.Next {
		if node.Value == value {
			return true
		}
	}
	return false
}

func (ll *DoublyLinkedList) PrintList() {
	node := ll.Head
	fmt.Printf("[")
	for node != nil {
		fmt.Printf(" %d ", node.Value)
		node = node.Next
	}
	fmt.Printf(" ]\n")
}
func (ll *DoublyLinkedList) PrintListFromTail() {
	node := ll.Tail
	fmt.Printf("[")
	for node != nil {
		fmt.Printf(" %d ", node.Value)
		node = node.Prev
	}
	fmt.Printf(" ]\n")
}

func (ll *DoublyLinkedList) FillList(values []int) {
	for _, value := range values {
		newNode := &Node{Value: value}
		if ll.Head == nil {
			ll.Head = newNode
			ll.Tail = newNode
		} else {
			ll.Tail.Next = newNode
			newNode.Prev = ll.Tail
			ll.Tail = newNode
		}
	}
}

func main() {
	linkedList := NewDoublyLinkedList()
	one := NewNode(1)
	two := NewNode(2)
	three := NewNode(3)
	four := NewNode(4)
	bindNodes(one, two)
	bindNodes(two, three)
	bindNodes(three, four)
	linkedList.Head = one
	linkedList.Tail = four
	linkedList.InsertAfter(one, two)
	linkedList.PrintList()
	linkedList.PrintListFromTail()
	linkedList.InsertAfter(two, three)
	linkedList.PrintList()
	linkedList.PrintListFromTail()
	linkedList.InsertAfter(three, four)
	linkedList.PrintList()
	linkedList.PrintListFromTail()
	linkedList.InsertAfter(two, one)
	linkedList.PrintList()
	linkedList.PrintListFromTail()
	linkedList.InsertBefore(three, four)
	linkedList.PrintList()
	linkedList.PrintListFromTail()
	fmt.Println("===================")
}
