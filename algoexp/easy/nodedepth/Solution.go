package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return NodeDepthsRec(root, 0, -1)
}

func NodeDepthsRec(root *BinaryTree, sum int, depth int) int {
	depth += 1
	sum += depth
	if root.Left == nil && root.Right == nil {
		return sum
	}
	if root.Left != nil {
		sum = NodeDepthsRec(root.Left, sum, depth)
	}
	if root.Right != nil {
		sum = NodeDepthsRec(root.Right, sum, depth)
	}
	return sum
}

func main() {
	root := &BinaryTree{1,
		&BinaryTree{2,
			&BinaryTree{4,
				&BinaryTree{8, nil, nil},
				&BinaryTree{9, nil, nil}},
			&BinaryTree{5, &BinaryTree{10, nil, nil}, nil}},
		&BinaryTree{3, &BinaryTree{6, nil, nil}, &BinaryTree{7, nil, nil}}}
	sum := NodeDepths(root)
	fmt.Print(sum)
}
