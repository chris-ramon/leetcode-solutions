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
var maxPathSum = function(root) {
    var result = Number.MIN_SAFE_INTEGER;
    if(!root) return result;
    
    const dfs = (node) => {
        if(!node) return 0;
        
        const leftPathSum = Math.max(dfs(node.left), 0);
        const rightPathSum = Math.max(dfs(node.right), 0);
        
        result = Math.max(result, leftPathSum + rightPathSum + node.val);
        
        var keptPathSum = Math.max(leftPathSum, rightPathSum);
        return node.val + keptPathSum;
    };
    
    dfs(root);
    
    return result;
};