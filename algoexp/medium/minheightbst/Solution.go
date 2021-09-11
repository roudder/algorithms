package main

func MinHeightBST(array []int) *BST {
	if len(array) < 1 {
		return nil
	}
	mid := len(array) / 2
	bst := BST{}
	bst.Value = array[mid]
	bst.Left = MinHeightBST(array[:mid])
	bst.Right = MinHeightBST(array[mid+1:])
	return &bst
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
