package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	switch {
	case tree.Left != nil && tree.Right != nil:
		tree.Right, tree.Left = tree.Left, tree.Right
		tree.Left.InvertBinaryTree()
		tree.Right.InvertBinaryTree()

	case tree.Left != nil && tree.Right == nil:
		tree.Right = tree.Left
		tree.Left = nil
		tree.Right.InvertBinaryTree()
	case tree.Right != nil && tree.Left == nil:
		tree.Left = tree.Right
		tree.Right = nil
		tree.Left.InvertBinaryTree()
	default:
		return
	}

}
