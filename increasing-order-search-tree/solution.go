/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    vals := inorder(root)
    sentinel := &TreeNode{}
    temp := sentinel
    
    for i := 0; i < len(vals); i++ {
        val := vals[i]
        node := TreeNode{Val: val}
        temp.Right = &node
        temp = temp.Right
    }
    
    return sentinel.Right
}

func inorder(node *TreeNode) []int {
    if node == nil {
        return []int{}
    }
    if node.Left == nil && node.Right == nil {
        return []int{node.Val}
    }
    
    leftVals := inorder(node.Left) 
    rightVals := inorder(node.Right) 
    
    vals := []int{}
    
    vals = append(vals, leftVals...)
    vals = append(vals, node.Val)
    vals = append(vals, rightVals...)
    
    return vals
}