/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    return equals(p, q)
}

func equals(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil && q != nil || p != nil && q == nil {
        return false
    }
    if p.Left == nil && p.Right == nil && q.Left == nil && q.Right == nil {
        return p.Val == q.Val
    }
    return p.Val == q.Val && equals(p.Left, q.Left) && equals(p.Right, q.Right)
}