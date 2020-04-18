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
var isValidBST = function(root) {
    const isValid = (node, min, max) => {
        if(!node) return true;
        if (min !== undefined && min <= node.val) {
            return false;
        }
        if (max !== undefined && max >= node.val) {
            return false;
        }
        return isValid(node.left, node.val, max) && isValid(node.right, min, node.val);
    };
    
    return isValid(root, undefined, undefined);
};

console.log(isValidBST({val: 2, left: {val: 1}, right: {val: 3}}));
