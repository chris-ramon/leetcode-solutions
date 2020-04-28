/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number}
 */
var maxDepth = function(root) {
    if(!root) return 0;
    var result = 0;
    var level = 1;
    const queue = [root];
    while(queue.length) {
        var size = queue.length;
        for(var i = 0; i < size; i++) {
            const node = queue.shift();
            if(!node.children || !node.children.length) {
                result = result ? Math.max(result, level) : level;
            } else {
                queue.push(...node.children);  
            }
        }
        level += 1;
    }
    return result;
};