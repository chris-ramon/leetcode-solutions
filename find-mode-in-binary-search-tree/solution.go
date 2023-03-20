/*
TestCases
- Time: O(n), where n is the size of the given tree.
- Space: O(n), where n is the size of the stack during the DFS.
*/

/*
TestCases
[1,null,2,2]\n[0]\n[1,null,3]\n
*/


func findMode(root *TreeNode) []int {
	countMap := map[int]int{}
	dfs(root, countMap)
	result := []int{}
	count := [][]int{}
	for k, v := range countMap {
		val := []int{k, v}
		count = append(count, val)
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i][1] > count[j][1]
	})
	for _, v := range count {
		if count[0][1] == v[1] {
			result = append(result, v[0])
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}

func dfs(node *TreeNode, countMap map[int]int) {
	if node == nil {
		return
	}
	countMap[node.Val]++
	dfs(node.Left, countMap)
	dfs(node.Right, countMap)
}
