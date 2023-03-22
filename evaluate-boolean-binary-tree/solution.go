/*
Runtime Complexity
- Time: O(n), where n is the size of the tree.
- Space: O(n), where n is the size of the callstack during dfs.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    return dfs(root)
}

func dfs(root *TreeNode) bool {
    if root == nil {
        return false
    }
    if root.Val == 0 {
        return false
    }
    if root.Val == 1 {
        return true
    }
    if root.Val == 2 {
        return dfs(root.Left) ||dfs(root.Right)
    }
    if root.Val == 3 {
        return dfs(root.Left) && dfs(root.Right)
    }
    return false
}
