package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	var node *BST
	node = &BST{value, nil, nil}

	if tree == nil {
		tree = node
	} else {
		insertTreeNode(tree, node)
	}
}

func insertTreeNode(tree *BST, newNode *BST) {
	if tree.Value < newNode.Value {
		if tree.Left == nil {
			tree.Left = newNode
		} else {
			insertTreeNode(tree.Left, newNode)
		}
	} else {
		insertTreeNode(tree.Right, newNode)
	}
	if tree.Value > newNode.Value {
		if tree.Right == nil {
			tree.Right = newNode
		} else {
			insertTreeNode(tree.Right, newNode)
		}
	}
}
