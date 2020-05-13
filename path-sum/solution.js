/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} sum
 * @return {boolean}
 */
var hasPathSum = function(root, sum) {
    if(!root) return false;
    const [acc, found] =  dfs(root, 0, sum);
    return found;
};

var dfs = function(node, acc, sum) {
    if(node == null) return [acc, false];
    if(node.right == null && node.left == null) {
        return [acc + node.val, acc + node.val == sum];
    }
    
    
    const [lAcc, lFound] = dfs(node.left, acc + node.val, sum);
    if(lFound) {
        return [lAcc, true];
    }
    
    const [rAcc, rFound] = dfs(node.right, acc + node.val, sum);
    if(rFound) {
        return [rAcc, true];
    }
    
    return [acc, false];
};