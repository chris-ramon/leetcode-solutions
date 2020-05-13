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
    var depth = new Map();
    var parents = new Map();
    dfs(root, null, x, y, depth, parents, 1);
    return depth[x] == depth[y] && parents[x] != parents[y];
};

var dfs = function(node, parent, x, y, depth, parents, level) {
    if(node == null) return;
    dfs(node.left, node, x, y, depth, parents, level + 1);
    dfs(node.right, node, x, y, depth, parents, level + 1);
    if(node.val === x) {
        depth[x] = level;
        parents[x] = parent;
    }
    if(node.val === y) {
        depth[y] = level;
        parents[y] = parent;
    }
};
    
