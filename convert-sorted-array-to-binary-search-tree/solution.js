/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var sortedArrayToBST = function(nums) {
    if(!nums || !nums.length) return null;
    return dfs(nums);
};

var dfs = function(nums) {
    if(!nums || !nums.length) return null;
    const mid = Math.floor(nums.length / 2);
    const node = new TreeNode(nums[mid]);
    node.left = dfs(nums.slice(0, mid), 0);
    node.right = dfs(nums.slice(mid + 1), 0);
    return node;
};