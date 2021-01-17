// Complexity Analysis
// Time Complexity: O(N), where N is the number of nodes in the binary tree. 
// Space Complexity: O(N), where N is the stack size during the DFS. 

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
 * @return {number}
 */

var findTilt = function(root) {
    var result = {sum: 0}
    
    dfs(root, result)
    
    return result.sum
};

var dfs = function(node, result) {
    if(node === null) {
        return 0
    }
    
    const leftSum = dfs(node.left, result)
    const rightSum = dfs(node.right, result)
    
    const absSum = Math.abs(leftSum - rightSum)
    
    result.sum += absSum
    
    return node.val + leftSum + rightSum
}


/*
Testcases:

[]
[1]
[-1,1]
[1,2,3]
[4,2,9,3,5,null,7]
[21,7,14,1,1,2,2,3,3]
*/ 