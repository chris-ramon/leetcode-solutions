/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} x
 * @param {number} y
 * @return {boolean}
 */
var isCousins = function(root, x, y) {
    if(!root) return false;
    var xDepth = null;
    var yDepth = null;
    var xParent = null;
    var yParent = null;
    var level = 1;
    var queue = [[root, null]];
    while(queue.length) {
        var size = queue.length;
        for(var i = 0; i < size; i++) {
            var [node, parent] = queue.shift();
            if(node.val == x) {
                xDepth = level;
                xParent = parent;
            }
            if(node.val == y) {
                yDepth = level;
                yParent = parent;
            }
            if(xDepth != null && yDepth != null) {
                return xDepth === yDepth && xParent !== yParent;
            }
            node.left && queue.push([node.left, node]);
            node.right && queue.push([node.right, node]);
        }
        level += 1;
    }
    return false;
};