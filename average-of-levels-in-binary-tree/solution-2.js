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
    const q = [root]
    const result = []
    while(q.length) {
        const size = q.length
        let sum = 0
        for(let i = 0; i < size; i++) {
            const node = q.shift()
            node.left && q.push(node.left)
            node.right && q.push(node.right)
            sum += node.val
        }
        result.push(sum / size)
    }
    return result
};