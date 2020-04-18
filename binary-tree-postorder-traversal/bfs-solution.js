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
    var queue = [root];
    while(queue.length > 0) {
        var size = queue.length;
        for (var i = 0; i < size; i++) {
            var node = queue.pop();
            if(!node) continue;
            result.push(node.val);
            queue.push(node.left);
            queue.push(node.right);
        }
    }
    return result.reverse();
};
