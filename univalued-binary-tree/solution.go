// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type result struct {
	value bool
 }

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	val := root.Val
	result := result{value: true}
	dfs(root, &result, val)
	return result.value
}

func dfs(node *TreeNode, result *result, val int) {
	if node == nil {
		return
	}
	if node.Val != val {
		result.value = false
		return
	}
	dfs(node.Left, result, val)
	dfs(node.Right, result, val)
}
// @lc code=end

/*
TestCases
[1,1,1,1,1,null,1]\n[2,2,2,5,2]\n
*/
