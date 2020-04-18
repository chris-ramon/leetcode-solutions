/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    // even - left to right, odd - right to left
    var result = [];
    if(!root) return result;
    var queue = [root];
    var level = 0;
    while(queue.length > 0) {
        var size = queue.length;
        var resultSubset = [];
        for (var i = 0; i < size; i++) {
            var node = queue.shift();
            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
            resultSubset.push(node.val);
        }
        if(level % 2 !== 0) {
            resultSubset = resultSubset.reverse();
        }
        result.push([...resultSubset]);
        resultSubset = [];
        level += 1;
    }
    return result;
};
