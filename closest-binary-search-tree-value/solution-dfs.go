// Complexity Analysis
// Time: `O(N)`, where N is the number of nodes of the tree.
// Space: `O(N)`, where N is the size of the stack during the depth-first search.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
    var result int
    var minDiff float64 = math.MaxFloat64
    dfs(root, target, &minDiff, &result)
    return result
}

func dfs(root *TreeNode, target float64, minDiff *float64, result *int) {
    if root == nil {
        return
    }
    dfs(root.Left, target, minDiff, result)
    dfs(root.Right, target, minDiff, result)
    diff := math.Abs(float64(root.Val) - target)
    if diff < *minDiff {
        *minDiff = diff
        *result = root.Val
    }
}

// Testcases:
/*
[4,2,5,1,3]
100000
[4,2,5,1,3]
4.4
[1]
1
*/
