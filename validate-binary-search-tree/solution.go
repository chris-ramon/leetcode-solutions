package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 30,
		},
	}
	log.Println(isValidBST(&tree))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return visitAndCheck(root, nil, nil)
}

func visitAndCheck(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val >= *min {
		return false
	}
	if max != nil && node.Val <= *max {
		return false
	}
	return visitAndCheck(node.Left, &node.Val, max) && visitAndCheck(node.Right, min, &node.Val)
}
