// Complexity Analysis
// Time: `O(M + N)`, where M is the number of nodes in t1 tree and N is the number of nodes in t2 tree.
// Space: `O(max(M,N))`, because the queue size can growp up to max(M,N), where M and N is the number of nodes of t1 and t2 tree respectively. 


/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} t1
 * @param {TreeNode} t2
 * @return {TreeNode}
 */
var mergeTrees = function(t1, t2) {
    if(!t1 && !t2) return null;
    
    var root = new TreeNode();
    var q = [[root, t1, t2]];
    
    while(q.length) {
        
        var size = q.length;
        
        for(var i = 0; i < size; i++) {
            var [n, n1, n2] = q.shift();
            if(!n) continue;
            
            n.val += (n1 && n1.val) || 0
            n.val += (n2 && n2.val) || 0
            
            if((n1 && n1.left) || (n2 && n2.left)) {
                n.left = new TreeNode(0);
            }
            
            if((n1 && n1.right) || (n2 && n2.right)) {
                n.right = new TreeNode(0);
            }
            
            q.push([n.left, (n1 && n1.left) || null, (n2 && n2.left) || null]);
            q.push([n.right, (n1 && n1.right) || null, (n2 && n2.right) || null]);
        }
    }
    
    return root;
};

// Testcases:
/*
[-1]
[]
[]
[1]
[-3]
[0]
[0]
[0]
[]
[]
[1]
[2]
[1,3,2]
[2,1,3]
[1,3,2,5]
[2,1,3,null,4,null,7]
*/