package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}
func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}

func inorderTraversal(root *BinaryTree) {
	printASCII(root, "", false)
}

func printASCII(tree *BinaryTree, prefix string, isLeft bool) {
	if tree == nil {
		return
	}

	if tree.Right != nil {
		printASCII(tree.Right, prefix+"       ", false)
	}
	fmt.Printf("%s", prefix)

	if isLeft {
		fmt.Printf("└──")
	} else {
		fmt.Printf("┌──")
	}
	fmt.Println(tree.Value)

	if tree.Left != nil {
		printASCII(tree.Left, prefix+"       ", true)
	}
}

func allPaths(root *BinaryTree) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Value}}
	}

	leftPaths := allPaths(root.Left)
	rightPaths := allPaths(root.Right)

	allPaths := make([][]int, 0)

	for _, path := range leftPaths {
		newPath := []int{root.Value}
		newPath = append(newPath, path...)
		allPaths = append(allPaths, newPath)
	}

	for _, path := range rightPaths {
		newPath := []int{root.Value}
		newPath = append(newPath, path...)
		allPaths = append(allPaths, newPath)
	}
	return allPaths
}

func DFS(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Value)
	DFS(root.Left)
	DFS(root.Right)
}

func branchSums(tree *BinaryTree, sum int, sums *[]int) {
	if tree == nil {
		return
	}

	sum += tree.Value
	if tree.Right == nil && tree.Left == nil {
		*sums = append(*sums, sum)
		return
	}
	branchSums(tree.Left, sum, sums)
	branchSums(tree.Right, sum, sums)
}

func BranchSums(tree *BinaryTree) []int {
	sums := []int{}
	branchSums(tree, 0, &sums)
	return sums
}

func main() {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	inorderTraversal(tree)
	fmt.Println(allPaths(tree))
	DFS(tree)
	fmt.Println(BranchSums(tree))
}
