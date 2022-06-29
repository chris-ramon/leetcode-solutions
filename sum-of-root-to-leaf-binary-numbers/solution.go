func sumRootToLeaf(root *TreeNode) int {
	paths := []string{}
	sum := 0.0
	dfs(root, "", &paths)
	for _, path := range paths {
		preSum := 0.0
		l := len(path)
		for i := l - 1; i >= 0; i-- {
			if path[i] == '1' {
				preSum += math.Pow(2, float64(l-1.0-i))
			}
		}
		sum += preSum
	}
	return int(sum)
}

func dfs(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}
	p := path + fmt.Sprintf("%d", node.Val)
	dfs(node.Left, p, paths)
	dfs(node.Right, p, paths)
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, p)
	}
}
