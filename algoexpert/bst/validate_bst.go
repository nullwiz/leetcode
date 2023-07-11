package main

import "fmt"

func (tree *BST) ValidateBst() bool {
	// Each node has at most, two children, this
	// Is validated by the default structure.
	// All values in the left subtree ar less than the nodes value
	if tree.Left != nil {
		if tree.Left.Value > tree.Value {
			return false
		}
		if !tree.Left.ValidateBst() {
			return false
		}
	}
	if tree.Right != nil {
		if tree.Right.Value < tree.Value {
			return false
		}
		if !tree.Right.ValidateBst() {
			return false
		}
	}
	return true
}

func createInvalidBST() *BST {
	var tree BST
	for _, v := range []int{10, 5, 15, 5, 2, 1, 22} {
		tree.Insert(v)
	}
	return &tree
}

func printBST(root *BST, level int, dir string) {
	if root == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "       "
	}
	format += "---[ "
	direction := dir
	format += direction
	level++
	printBST(root.Left, level, "L:")
	printBST(root.Right, level, "R:")

	fmt.Printf(format+"%d\n", root.Value)
}

func main_validate() {
	// Validate BST
	// Write a function that takes in a potentially invalid Binary Search Tree (BST)
	// and returns a boolean representing whether the BST is valid.
	var BST = createInvalidBST()
	printBST(BST, 0, "root")
	if BST.ValidateBst() {
		fmt.Println("BST is valid")
	} else {
		fmt.Println("BST is invalid")
	}
}
