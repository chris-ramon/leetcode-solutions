// Complexity Analysis
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
func pathSum(root *TreeNode, sum int) int {
    var queue []*TreeNode = []*TreeNode{root}
    var vals map[*TreeNode][]int = make(map[*TreeNode][]int, 0)
    var count int
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            if node == nil {
                continue
            }
            nVals := vals[node]
            if node.Val == sum {
                count += 1
            }
            var partialSum int
            for i := len(nVals) - 1; i >= 0; i-- {
                partialSum += nVals[i]
                if (node.Val + partialSum) == sum {
                    count += 1
                }
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
                vals[node.Left] = []int{}
                vals[node.Left] = append(vals[node.Left], nVals...)
                vals[node.Left] = append(vals[node.Left], node.Val)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
                vals[node.Right] = []int{}
                vals[node.Right] = append(vals[node.Right], nVals...)
                vals[node.Right] = append(vals[node.Right], node.Val)
            }
        }
    }
    return count
}
