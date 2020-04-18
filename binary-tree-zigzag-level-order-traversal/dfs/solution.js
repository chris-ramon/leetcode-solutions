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
    const result = [];
    if (!root) return result;
    const results = {};
    const dfs = (node, level) => {
        if (!node) return;
        if (results[level] !== undefined) results[level].push(node.val);
        if (results[level] === undefined) results[level] = [node.val];
        dfs(node.left, level + 1);
        dfs(node.right, level + 1);
    };
    dfs(root, 0);
    for (var k in results) {
        var resultSubset = results[k];
        if (k % 2 !== 0) resultSubset.reverse();
        result.push([...resultSubset]);
    }
    return result;
};
