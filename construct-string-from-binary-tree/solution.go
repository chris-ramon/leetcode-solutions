/*
Complexity Analysis
- Time: O(n), where n is the size of given tree.
- Space: O(n), where n is the size of given tree.
*/

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
	result := ""
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		if _, found := visited[n]; found {
			stack = stack[:len(stack)-1]
			result += ")"
		} else {
			visited[n] = true
			result += fmt.Sprintf("(%d", n.Val)
			if n.Left == nil && n.Right != nil {
				result += "()"
			}
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
		}
	}
	return result[1 : len(result)-1]
}
