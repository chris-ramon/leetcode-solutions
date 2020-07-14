// Complexity Analysis:
// Time: `O(N)`, where N is the number of nodes in the tree.
// Space: `O(N)`, where N is the size of the stack during the depth-first search.
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    var result int
    dfs(root, L, R, &result)
    return result
}

func dfs(root *TreeNode, L int, R int, result *int) {
    if root == nil {
        return
    }
    dfs(root.Left, L, R, result)
    dfs(root.Right, L, R, result)
    if root.Val >= L && root.Val <= R {
        *result += root.Val
    }
}

// Testcases:
/*
[10,5,15,3,7,13,18,1,null,6]
6
10
[10,5,15,3,7,null,18]
7
15
[1]
1
1
[]
0
0
*/
