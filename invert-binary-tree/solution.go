// Complexity Analysis
// Time: `O(N)`, where N is the number of nodes in the tree.
// Space: `O(N)`, where N is the size of the call stack during the depth-first search.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    head := root
    dfs(head)
    return head
}

func dfs(root *TreeNode) {
    if root == nil {
        return
    }
    
    l := root.Left
    r := root.Right
    
    root.Left = r
    root.Right = l
    
    dfs(root.Left)
    dfs(root.Right)
}

// Testcases:
/*
[]
[4]
[4,2,7]
[4,2,7,1,3,6,9]
[16,8,28,4,12,24,36,1,5]
*/