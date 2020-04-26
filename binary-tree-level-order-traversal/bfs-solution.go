/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    
    var level int = 1
    var results map[int][]int = make(map[int][]int, 0)
    
    dfs(root, results, level)
    
    var levels int = len(results)
    for i := 1; i <= levels; i++ {
        values := results[i]
        result = append(result, values)
    }
    
    return result
}

func dfs(node *TreeNode, results map[int][]int, level int) {
    if node == nil {
        return
    }
    
    if _, found := results[level]; found {
        results[level] = append(results[level], node.Val)
    } else {
        results[level] = []int{node.Val}
    }
    
    dfs(node.Left, results, level + 1)
    dfs(node.Right, results, level + 1)
}