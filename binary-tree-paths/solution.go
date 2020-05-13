/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }
    return paths(root, "")
}

func path(p string, v int) string {
    if p == "" {
        return fmt.Sprintf("%d", v)
    }
    return fmt.Sprintf("%s->%d", p, v)
}

func paths(node *TreeNode, currentPath string) []string {
    if node == nil {
        return []string{}
    }
    
    if node.Left == nil && node.Right == nil {
        return []string{path(currentPath, node.Val)}
    }
    
    leftPaths := paths(node.Left, path(currentPath, node.Val))
    rightPaths := paths(node.Right, path(currentPath, node.Val))
    
    var result []string
    result = append(result, leftPaths...)
    result = append(result, rightPaths...)
    
    return result
}