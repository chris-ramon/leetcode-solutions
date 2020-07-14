/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
    var count int
    dfsStarter(root, &count, sum)
    return count
}

func dfsStarter(root *TreeNode, count *int, sum int) {
    if root == nil {
        return
    }
    dfs(root, 0, count, sum)
    dfsStarter(root.Left, count, sum)
    dfsStarter(root.Right, count, sum)
}

func dfs(root *TreeNode, currentSum int, count *int, sum int) {
    if root == nil {
        return
    }
    dfs(root.Left, currentSum + root.Val, count, sum)
    dfs(root.Right, currentSum + root.Val, count, sum)
    if root.Val + currentSum == sum {
        *count += 1
    }
}

// Testcases:
/*
[1]
1
[2,3]
2
[1,2]
2
[10,5,-3,3,2,null,11,3,-2,null,1]
8
[1,1,1,1,1,1,1,1]
2
[]
0
[1,-2,-3,1,3,-2,null,-1]
-1
[1,0,1,1,2,0,-1,0,1,-1,0,-1,0,1,0]
2
[10,5,5,10,10,10,5,null,null,5,null,null,null,null,null,5,null,5]
15
[10,5,-3,3,2,null,11,3,-2,null,1]
8
[5,4,8,11,null,13,4,7,2,null,null,5,1]
22
*/
