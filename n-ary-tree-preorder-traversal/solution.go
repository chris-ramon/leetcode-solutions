/*
Complexity Analysis
- Time: O(n), where n is the size of the given tree.
- Space: O(n), where n is the size of the call stack during the dfs.
*/


func preorder(root *Node) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	for _, n := range node.Children {
		dfs(n, result)
	}
}
