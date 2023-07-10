/*
Complexity Analysis:

- Time: O(n), where n is the size of the tree.
- Space: O(n), where n is the size of the call-stack during the DFS.
*/


/**
 * Definition for a rope tree node.
 * type RopeTreeNode struct {
 * 	   len   int
 * 	   val   string
 * 	   left  *RopeTreeNode
 * 	   right *RopeTreeNode
 * }
 */
func getKthCharacter(root *RopeTreeNode, k int) byte {
	str := ""
	dfs(root, &str)
	return str[k-1]
}

func dfs(node *RopeTreeNode, str *string) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		*str += node.val
	}
	dfs(node.left, str)
	dfs(node.right, str)
}
