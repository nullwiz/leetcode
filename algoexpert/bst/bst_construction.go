package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if tree == nil {
		return &BST{Value: value}
	} else if value < tree.Value {
		tree.Left = tree.Left.Insert(value)
	} else {
		tree.Right = tree.Right.Insert(value)
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	if tree == nil {
		return false
	} else if value < tree.Value {
		return tree.Left.Contains(value)
	} else if value > tree.Value {
		return tree.Right.Contains(value)
	}
	return true
}

func (tree *BST) Remove(value int) *BST {
	if tree == nil {
		return nil
	}

	if value < tree.Value {
		tree.Left = tree.Left.Remove(value)
	} else if value > tree.Value {
		tree.Right = tree.Right.Remove(value)
	} else {
		if tree.Left == nil && tree.Right == nil {
			tree = nil
		} else if tree.Right != nil {
			leftmost := tree.Right
			for {
				if leftmost.Left == nil {
					break
				}
				leftmost = leftmost.Left
			}
			valueToReplace := leftmost.Value
			tree.Value = valueToReplace
			tree.Right = tree.Right.Remove(valueToReplace)
		} else {
			rightmost := tree.Left
			for {
				if rightmost.Right == nil {
					break
				}
				rightmost = rightmost.Right
			}
			valueToReplace := rightmost.Value
			tree.Value = valueToReplace
			tree.Left = tree.Left.Remove(valueToReplace)
		}
	}
	return tree
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}

func (tree *BST) getMaxValue() int {
	if tree.Right == nil {
		return tree.Value
	}
	return tree.Right.getMaxValue()
}

func main() {
	main_validate()
}
