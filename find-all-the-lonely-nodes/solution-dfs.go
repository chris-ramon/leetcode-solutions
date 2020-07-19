// Complexity Analysis
// Time: `O(N)`, where N is the number of nodes of the tree.
// Space: `O(N)`, where N is the stack size during the depth-first search.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
    var result []int
    dfs(root, false, &result)
    return result
}

func dfs(root *TreeNode, isLonely bool, result *[]int) {
    if root == nil {
        return
    }
    if isLonely {
        *result = append(*result, root.Val)
    }
    var childs int
    if root.Left != nil {
        childs += 1
    }
    if root.Right != nil {
        childs += 1
    }
    dfs(root.Left, childs == 1, result)
    dfs(root.Right, childs == 1, result)
}

// Testcases:
/*
[1,2,3,null,4]
[7,1,4,6,null,5,3,null,null,null,null,null,2]
[11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
[197]
[31,null,78,null,28]
[1]
*/