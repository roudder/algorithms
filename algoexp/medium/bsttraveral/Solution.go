package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array)
	}
	array = append(array, tree.Value)

	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array)
	}

	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	array = append(array, tree.Value)

	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}

	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.

	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}

	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)

	return array
}

func main() {
	bst10 := BST{}
	bst10.Value = 10
	bst5 := BST{}
	bst5.Value = 5
	bst2 := BST{}
	bst2.Value = 2
	bst1 := BST{}
	bst1.Value = 1
	bst55 := BST{}
	bst55.Value = 5
	bst15 := BST{}
	bst15.Value = 15
	bst13 := BST{}
	bst13.Value = 13
	bst22 := BST{}
	bst22.Value = 22
	bst14 := BST{}
	bst14.Value = 14

	bst10.Left = &bst5
	bst10.Right = &bst15
	bst5.Left = &bst2
	bst5.Right = &bst55
	bst2.Left = &bst1

	bst15.Left = &bst13
	bst15.Right = &bst22
	bst13.Right = &bst14

	traversed := bst10.InOrderTraverse([]int{})
	pre := bst10.PreOrderTraverse([]int{})
	post := bst10.PostOrderTraverse([]int{})
	fmt.Println(traversed)
	fmt.Println(pre)
	fmt.Println(post)
}
