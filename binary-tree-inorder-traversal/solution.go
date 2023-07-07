/*
Complexity Analysis:
- Time: O(n), where n is the size of the tree.
- Space: O(n), where n is the size of the call-stack during the DFS.
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    dfs(root, &result)
    return result
}

func dfs(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    dfs(node.Left, result)
    *result = append(*result, node.Val)
    dfs(node.Right, result)
}
