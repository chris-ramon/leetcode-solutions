// https://leetcode.com/problems/binary-tree-postorder-traversal/
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var postorderTraversal = function(root) {
    var result = [];
    if (!root) return result;
    const dfs = (node) => {
        if(!node) return;
        dfs(node.left);
        dfs(node.right);
        result.push(node.val);
    };
    dfs(root);
    return result;
};
