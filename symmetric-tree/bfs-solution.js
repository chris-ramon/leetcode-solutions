/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSymmetric = function(root) {
    var result = true;
    if(!root) return result;
    const queue = [root];
    const nodesValByLevel = {};
    var level = 0;
    while(queue.length) {
        const size = queue.length;
        level += 1;
        for (var i = 0; i < size; i++) {
            var node = queue.shift();
            var val = "";
            if (!node) {
                 val = "#"
            } else {
                val = node.val;
            }
            if (nodesValByLevel[level] === undefined) {
                nodesValByLevel[level] = [val];
            } else {
                nodesValByLevel[level].push(val);
            }
            if (!node) continue;
            queue.push(node.left, node.right);
        }
    }
    for (var level in nodesValByLevel) {
        const values = nodesValByLevel[level];
        const size = values.length;
        const left = values.slice(0, size/2);
        const right = values.slice(size/2).reverse();
        if (values.length > 1 && left.join("") !== right.join("")) {
            result = false;
            break;
        }
    }
    return result;
};