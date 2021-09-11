package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Validate() bool {
	return tree.ValidateCustom(math.MinInt32, math.MaxInt32)
}

func (tree *BST) ValidateCustom(minValue, maxValue int) bool {
	if tree == nil {
		return true
	}
	if tree.Value < minValue || tree.Value >= maxValue {
		return false
	}
	leftValid := tree.Left.ValidateCustom(minValue, tree.Value)
	return leftValid && tree.Right.ValidateCustom(tree.Value, maxValue)
}

func main() {

}
