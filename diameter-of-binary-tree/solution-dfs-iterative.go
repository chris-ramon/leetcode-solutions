/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var stack []*TreeNode = []*TreeNode{root}
    var visited map[*TreeNode]bool = make(map[*TreeNode]bool, 0)
    var depths map[*TreeNode]float64 = make(map[*TreeNode]float64, 0)
    var result float64
    for len(stack) > 0 {
        size := len(stack)
        for i := 0; i < size; i++ {
            node := stack[len(stack)-1]
            visited[node] = true
            if node == nil ||(node.Left == nil && node.Right == nil) {
                stack = stack[:len(stack)-1]
                depths[node] = 1
                continue
            }
            if node.Right != nil && !visited[node.Right] {
                stack = append(stack, node.Right)
            }
            if node.Left != nil && !visited[node.Left] {
                stack = append(stack, node.Left)
            }
            if (node.Right != nil && visited[node.Right]) || (node.Left != nil && visited[node.Left]) {
                l := depths[node.Left]
                r := depths[node.Right]
                pathLen := l + r
                result = math.Max(result, pathLen)
                depths[node] = math.Max(l, r) + 1
                stack = stack[:len(stack)-1]
            }
        }
    }
    return int(result)
}

// Testcases:
/*
[]
[1]
[1,2]
[1,2,3]
[1,2,3,4,5]
[1,2,3,4,null,null,null,null,5,6,7]
[1,2,3,4,null,null,null,null,5,6,7,null,null,null,10]
*/
