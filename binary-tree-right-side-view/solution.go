/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var result []int
    
    if root == nil {
        return result 
    }
    
    var (
        levelsMap map[int][]*TreeNode = make(map[int][]*TreeNode, 0)   
        level int = 0
        levels []int
        queue []*TreeNode = []*TreeNode{root}
    )
    
    level += 1
    levels = append(levels, level)
    
    for len(queue) > 0 {
        var size = len(queue)
        
        level += 1
        levels = append(levels, level)
        
        for i := 0; i < size; i++ {
            treeNode := queue[0]
            queue = queue[1:]

            if treeNode.Left != nil {
                queue = append(queue, treeNode.Left)
                if _, found := levelsMap[level]; found {
                    levelsMap[level] = append(levelsMap[level], treeNode.Left)
                } else {
                    levelsMap[level] = []*TreeNode{treeNode.Left}
                }
            }
            
            if treeNode.Right != nil {
                queue = append(queue, treeNode.Right)
                if _, found := levelsMap[level]; found {
                    levelsMap[level] = append(levelsMap[level], treeNode.Right)
                } else {
                    levelsMap[level] = []*TreeNode{treeNode.Right}
                }
            }
        }
    }
    
    levelsMap[1] = []*TreeNode{root}
    
    for _, l := range levels {
        treeNodes := levelsMap[l]
        if(len(treeNodes) == 0) {
            continue
        }
        mostRightTreeNode := treeNodes[len(treeNodes)-1]
        result = append(result, mostRightTreeNode.Val)
    }
    
    return result
}