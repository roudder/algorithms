package main

import "fmt"

// BST Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	var bst *BST
	if value < tree.Value {
		if tree.Left == nil {
			bst = &BST{}
			bst.Value = value
			tree.Left = bst
		} else {
			bst = tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			bst = &BST{}
			bst.Value = value
			tree.Right = bst
		} else {
			bst = tree.Right.Insert(value)
		}
	}
	return bst
}

func (tree *BST) Contains(value int) bool {
	if value < tree.Value && tree.Left != nil {
		return tree.Left.Contains(value)
	} else if value > tree.Value && tree.Right != nil {
		return tree.Right.Contains(value)
	} else {
		if value == tree.Value {
			return true
		} else {
			return false
		}
	}
}

func (tree *BST) Remove(value int) *BST {
	tree.RemoveCustom(value, nil)
	return tree
}

func (tree *BST) RemoveCustom(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.RemoveCustom(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.RemoveCustom(value, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.RemoveCustom(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			} else {

			}

		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Left != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}

	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
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

	fmt.Println(bst10.Contains(19))
	insert := bst10.Insert(12)
	fmt.Println(insert)
	fmt.Println(bst10.Insert(11))
}
