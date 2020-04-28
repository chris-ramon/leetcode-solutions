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
var levelOrderBottom = function(root) {
    var result = [];
    
    if(root == null) return result;
    
    var queue = [root];
    var resultSubset = [];
    
    while(queue.length > 0) {
        var size = queue.length;
        
        for(var i = 0; i < size; i++) {
            var node = queue.shift();
            
            resultSubset.push(node.val);
            
            if(i === size - 1) {
                result.unshift(resultSubset);
            }
            
            if(node.left != null) queue.push(node.left);
            if(node.right != null) queue.push(node.right);
        }
        
        resultSubset = [];
    }
    
    return result;
};