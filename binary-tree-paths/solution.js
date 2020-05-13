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
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    if(!root) return [];
    return paths(root, "");
};

var paths = function(node, path) {
    if(!node) return "";
    if(node.left == null && node.right == null) {
        return [path + node.val];
    }
    var leftPaths = paths(node.left, path + node.val + "->");
    var rightPaths = paths(node.right, path + node.val + "->");
    return [...leftPaths, ...rightPaths];
};