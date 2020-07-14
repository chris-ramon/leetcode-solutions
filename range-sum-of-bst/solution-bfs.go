// Complexity Analysis:
// Time: `O(N)`, where N is the number of nodes in the tree.
// Space: `O(N)`, where N is the size of the queue during the breadth-first search.
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
    if root == nil {
        return result
    }
    var queue []*TreeNode = []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Val >= L && node.Val <= R {
                result += node.Val
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return result
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
