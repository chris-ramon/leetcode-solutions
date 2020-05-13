/**
 
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var increasingBST = function(root) {
    if(!root) return null;
    const nodes = rearrange(root);
    var head =  nodes[0];
    head.left = null;
    head.right = null;
    var tmp = head;
    for (var i = 1; i < nodes.length; i++) {
        nodes[i].left = null;
        nodes[i].right = null;
        
        tmp.right = nodes[i];
        
        tmp = tmp.right;
    }
    return head;
};

var rearrange = function(node) {
    if(node == null) return [];
    if(node.left == null && node.right == null) {
        return [node];
    }
    var leftNodes = rearrange(node.left);
    var rightNodes = rearrange(node.right);
    return [...leftNodes, node, ...rightNodes];
};