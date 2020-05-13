/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    _, isBalanced := checkBalance(root, 0)
    return isBalanced
}

func checkBalance(node *TreeNode, level int) (int, bool) {
    if node == nil {
        return level, true 
    }
    
    if node.Left == nil && node.Right == nil {
        return level + 1, true
    }
    
    leftHeight, isLeftBalanced := checkBalance(node.Left, level + 1) 
    rightHeight, isRightBalanced := checkBalance(node.Right, level + 1) 
    
    if !isLeftBalanced {
        return leftHeight, false
    }
    
    if !isRightBalanced {
        return leftHeight, false
    }
    
    if leftHeight == rightHeight {
        return leftHeight, true
    }
    
    if leftHeight > rightHeight &&  leftHeight - rightHeight <= 1 {
        return leftHeight, true
    }
    
    if rightHeight > leftHeight &&  rightHeight - leftHeight <= 1 {
        return rightHeight, true
    }
    return leftHeight, false
}