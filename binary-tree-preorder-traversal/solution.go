// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
TestCases
- Time: O(n), where n is the number of nodes in the given tree.
- Space: O(n), where n is the size of the stack during the DFS.
*/

var result = []int{}

func preorderTraversal(root *TreeNode) []int {
	result = []int{}
	dfs(root)
	return result
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
}
// @lc code=end


/*
TestCases
[1,null,2,3]\n[2,3,4,5,6,7]\n[]\n[1]\n
*/
