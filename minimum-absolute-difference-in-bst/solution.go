/*
Complexity Analysis
- Time: O(n), where n is the size of the given tree.
- Space: O(n), where n is queue used to store the nodes during the BFS.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	q := []*TreeNode{root}
	values := []int{}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			values = append(values, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	result := float64(math.MaxInt)
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			v1 := float64(values[i])
			v2 := float64(values[j])
			if v1 != v2 {
				result = math.Min(result, math.Abs(v1 - v2))
			}
		}
	}
	return int(result)
}
