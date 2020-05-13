/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    return leafSequence(root1) == leafSequence(root2)
}

func leafSequence(node *TreeNode) string {
    if node == nil {
        return ""
    }
    if node.Left == nil && node.Right == nil {
        return fmt.Sprintf("-%d", node.Val)
    }
    leftSeq := leafSequence(node.Left)
    rightSeq := leafSequence(node.Right)
    return fmt.Sprintf("%s%s", leftSeq, rightSeq)
}