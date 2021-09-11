package main

import "fmt"

// Time comp.: 0(n*log(n)) Memory comp.: 0(1)
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var branchValues []int
	var sum int
	return BranchSumsRecursive(root, branchValues, sum)
}

func BranchSumsRecursive(root *BinaryTree, branchValues []int, sum int) []int {
	sum += root.Value

	if root.Left == nil && root.Right == nil {
		branchValues = append(branchValues, sum)
	}

	if root.Left != nil {
		branchValues = BranchSumsRecursive(root.Left, branchValues, sum)
	}
	if root.Right != nil {
		branchValues = BranchSumsRecursive(root.Right, branchValues, sum)
	}

	return branchValues
}

func main() {
	root := &BinaryTree{1,
		&BinaryTree{2,
			&BinaryTree{4,
				&BinaryTree{8, nil, nil},
				&BinaryTree{9, nil, nil}},
			&BinaryTree{5, &BinaryTree{10, nil, nil}, nil}},
		&BinaryTree{3, &BinaryTree{6, nil, nil}, &BinaryTree{7, nil, nil}}}
	branchValues := BranchSums(root)
	fmt.Print(branchValues)

}
