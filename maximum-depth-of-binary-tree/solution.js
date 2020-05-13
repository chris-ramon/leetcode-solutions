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
 * @return {number}
 */
var maxDepth = function(root) {
    if(!root) return 0;
    return depth(root, 0);
};

var depth = function(node, level) {
    if(node == null) return level;
    return Math.max(depth(node.left, level + 1), depth(node.right, level + 1))
};