/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    results := make(map[int][]int, 0)
    dfs(root, 0, results)
    levelsSorted := []int{}
    for level, _ := range results {
        levelsSorted = append(levelsSorted, level)
    }
    sort.Sort(sort.IntSlice(levelsSorted))
    for level := range levelsSorted {
        values, _ := results[level]
        if level % 2 != 0 {
            /*
            var valuesReversed []int 
            for r := len(values) - 1; r >= 0; r-- {
                v := values[r]
                valuesReversed = append(valuesReversed, v)
            }
            result = append(result, valuesReversed)
            */
            for r := len(values)/2-1; r >= 0; r-- {
                idx := len(values)-1-r
                values[r], values[idx] = values[idx], values[r]
            }
        }
        result = append(result, values)
    }
    return result
}

func dfs(node *TreeNode, level int, results map[int][]int) {
    if node == nil {
        return
    }
    if _, found := results[level]; found {
        results[level] = append(results[level], node.Val)
    } else {
        results[level] = []int{node.Val}
    }
    dfs(node.Left, level + 1, results)
    dfs(node.Right, level + 1, results)
}



