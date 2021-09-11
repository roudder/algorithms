package main

import (
	"fmt"
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	var diff, closest int
	if target == tree.Value {
		return tree.Value
	} else {
		closest = tree.Value
		diff = int(math.Abs(float64(target - tree.Value)))
	}
	if target < tree.Value && tree.Left != nil {
		left := tree.Left.FindClosestValue(target)
		if left == target {
			return left
		}
		leftDiff := int(math.Abs(float64(target - left)))
		if leftDiff < diff {
			closest = left
			diff = leftDiff
		}
	}
	if target > tree.Value && tree.Right != nil {
		right := tree.Right.FindClosestValue(target)
		if right == target {
			return right
		}
		rightDiff := int(math.Abs(float64(target - right)))
		if rightDiff < diff {
			closest = right
			diff = rightDiff
		}
	}
	return closest
}

func main() {
	bstL := &BST{Value: 50000, Left: nil, Right: nil}
	bstR := &BST{Value: 4500, Left: nil, Right: nil}
	bst := &BST{Value: 15, Left: bstL, Right: bstR}
	fmt.Print(bst.FindClosestValue(4500))
}
