/*
Complexity Analysis:

- Time complexity: O(n), where n is the size of the N-ary Tree.
- Space complexity: O(n), where n is the size of the call-stack during the DFS.
*/


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *Node, result *[]int) {
	if node == nil {
		return
	}
	for _, n := range node.Children {
		dfs(n, result)
	}
	*result = append(*result, node.Val)
}
