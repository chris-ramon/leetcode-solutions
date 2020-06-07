/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} s
 * @param {TreeNode} t
 * @return {boolean}
 */
var isSubtree = function(s, t) {
    var queue = [];
    queue.push(s);
    while(queue.length) {
        var node = queue.shift();
        if(node === null) continue;
        if(node.val === t.val && isEqual(node, t)) {
            return true;
        }
        queue.push(node.left);
        queue.push(node.right);
    }
    return false;
};

var isEqual = function(node1, node2) {
    var queue = [];
    queue.push({n1: node1, n2: node2});
    while(queue.length) {
        var {n1, n2} = queue.shift();
        if(n1 === null && n2 === null) continue;
        if(n1 === null || n2 === null || n1.val !== n2.val) {
            return false;
        }
        queue.push({n1: n1.left, n2: n2.left});
        queue.push({n1: n1.right, n2: n2.right});
    }
    return true;
};