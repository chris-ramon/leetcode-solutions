/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var averageOfLevels = function(root) {
    let q = [root]
    const result = []
    while(q.length) {
        const temp = []
        let sum = 0
        let size = q.length
        while(q.length) {
            const node = q.pop()
            node.left && temp.push(node.left)
            node.right && temp.push(node.right)
            sum += node.val
        }
        result.push(sum / size)
        q = temp
    }
    return result
};