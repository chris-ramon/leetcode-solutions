/*
Complexity Analysis
- Time: O(n), where n is the size of the tree.
- Space: O(n), where n is the size of the queue.
*/

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			continue
		}

		q = append(q, node.Left)
		q = append(q, node.Right)
		result++
	}

	return result
}

// @lc code=end
