/*
**Complexity Analysis**
Time: `O(N)`, where N is the number of nodes for the given tree, which are visited once.
Space: `O(N)`, where N is the size of the stack during the depth-first search.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var result float64
    dfs(root, &result)
    return int(result)
}

func dfs(root *TreeNode, result *float64) float64 {
    if root == nil {
        return 0
    }
    l := dfs(root.Left, result)
    r := dfs(root.Right, result)
    *result = math.Max(*result, l + r)
    return math.Max(l, r) + 1
}

// Testcases:
/*
[1]
[1,2]
[1,2,3]
[1,2,3,4,5]
[1,2,3,4,null,null,null,null,5,6,7]
[1,2,3,4,null,null,null,null,5,6,7,null,null,null,10]
*/
