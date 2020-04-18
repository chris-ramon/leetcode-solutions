package main

func main() {
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var result [][]int = [][]int{}
    if root == nil {
        return result
    }
    var queue []*TreeNode = []*TreeNode{root}
    var level int = 0
    for len(queue) > 0 {
        var size int = len(queue)
        var resultSubset []int
        for  i := 0; i < size; i++ {
            var node *TreeNode = queue[0];
            queue = queue[1:];
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            resultSubset = append(resultSubset, node.Val)
        }
        if level % 2 != 0 {
            /*
            var resultSubsetReversed []int
            for r := len(resultSubset) - 1; r >= 0; r-- {
                resultSubsetReversed = append(resultSubsetReversed, resultSubset[r])
            }
            // resultSubset = resultSubsetReversed
            */
            for r := len(resultSubset)/2 - 1; r >= 0; r-- {
                idx := len(resultSubset) - 1 - r
                resultSubset[r], resultSubset[idx] = resultSubset[idx], resultSubset[r]
            }
        }
        result = append(result, resultSubset)
        resultSubset = []int{}
        level += 1
    }
    return result
}
