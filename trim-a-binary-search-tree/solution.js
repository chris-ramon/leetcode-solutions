// Complexity Analysis
// Time: `O(N)`, where N is the number of nodes in the tree.
// Space: `O(N)`, where N is the callstack size during the depth-first search.

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
 * @param {number} L
 * @param {number} R
 * @return {TreeNode}
 */
var trimBST = function(root, L, R) {
    return dfs(root, L, R);
};

var dfs = function(root, L, R) {
    if(root == null) {
        return null;
    }
    if(root.val < L) {
        return dfs(root.right, L, R);
    }
    if(root.val > R) {
        return dfs(root.left, L, R);
    }
    root.left = dfs(root.left, L, R);
    root.right = dfs(root.right, L, R);
    return root;
}

// Testcases:
/*
[1]
2
2
[1,0,2]
1
2
[3,0,4,null,2,1]
1
3
*/