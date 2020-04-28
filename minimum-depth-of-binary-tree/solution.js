/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var minDepth = function(root) {
    if(root == null) return 0;
    
    var level = 1;
    var minDepth = 0;
    var queue = [root];
    
    while(queue.length) {
        var size = queue.length;
        for(var i = 0; i < size; i++) {
            var node = queue.shift();
            if(node.left == null && node.right == null) {
                minDepth = minDepth > 0 ? Math.min(minDepth, level): level;
            }
            if(node.left != null) queue.push(node.left);
            if(node.right != null) queue.push(node.right);
        }
        level += 1;
    }
    
    return minDepth;
};