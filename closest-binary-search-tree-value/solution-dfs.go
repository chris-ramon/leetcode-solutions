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
[41,37,44,24,39,42,48,1,35,38,40,null,43,46,49,0,2,30,36,null,null,null,null,null,null,45,47,null,null,null,null,null,4,29,32,null,null,null,null,null,null,3,9,26,null,31,34,null,null,7,11,25,27,null,null,33,null,6,8,10,16,null,null,null,28,null,null,5,null,null,null,null,null,15,19,null,null,null,null,12,null,18,20,null,13,17,null,null,22,null,14,null,null,21,23]
3.285714
*/
