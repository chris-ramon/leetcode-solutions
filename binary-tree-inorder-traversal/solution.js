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
var inorderTraversal = function(root) {
    const result = []
    if(!root) return result
    dfs(root, result)
    return result
};

const dfs = (node, result) => {
    if(!node) {
        return
    }
    dfs(node.left, result)
    result.push(node.val)
    dfs(node.right, result)
}