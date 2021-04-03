package main

// Complexity Analysis
// Time: O(N), where N is the number of nodes in the tree.
// Space: O(N), where N is the callstack size during the DFS.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	var (
		firstMin  float64 = math.MaxFloat64
		secondMin float64 = math.MaxFloat64
	)

	dfs(root, &firstMin, &secondMin)

	if firstMin != secondMin && secondMin != math.MaxFloat64 {
		return int(secondMin)
	}

	return -1
}

func dfs(root *TreeNode, firstMin *float64, secondMin *float64) {
	if root == nil {
		return
	}

	v := float64(root.Val)

	*firstMin = math.Min(*firstMin, v)

	if v > *firstMin {
		*secondMin = math.Min(*secondMin, v)
	}

	dfs(root.Left, firstMin, secondMin)
	dfs(root.Right, firstMin, secondMin)
}

/*
Testcases:

[2,2,5,null,null,5,7]


[2,2,2]
*/
