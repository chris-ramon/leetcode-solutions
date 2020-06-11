/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
    var queue = []; var pFound = false; var qFound = false; var result = null;
    var pAncestors = []; var qAncestors = [];
    queue.push({n: root, ancestors: []});
    while(queue.length) {
        var {n, ancestors} = queue.shift();
        if(n === null) continue;
        
        var _ancestors = [...ancestors, n];
        
        if(n.val === p.val) {
            pFound = true;
            pAncestors = _ancestors;
        }
        if(n.val === q.val) {
            pFound = true;
            qAncestors = _ancestors;
        } 
        
        if(pFound && qFound) break;
        
        queue.push({n: n.left, ancestors: _ancestors});
        queue.push({n: n.right, ancestors: _ancestors});
    }
    
    var qAncestorsMap = new Map();
    
    qAncestors.forEach(qAncestor => qAncestorsMap.set(qAncestor.val, qAncestor));
    
    for(var i = pAncestors.length - 1; i >= 0; i--) {
        var pAncestor = pAncestors[i];
        if(qAncestorsMap.has(pAncestor.val)) {
            return pAncestor;
        }
    }
    
    return result;
};