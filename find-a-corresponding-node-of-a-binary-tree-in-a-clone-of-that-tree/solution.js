// **Complexity Analysis**
// - Time: O(N) where N is the number of nodes in the binary tree.
// - Space: O(N) where N is the size of the queues during the breadth first search.

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} original
 * @param {TreeNode} cloned
 * @param {TreeNode} target
 * @return {TreeNode}
 */

var getTargetCopy = function(original, cloned, target) {
    q1 = [original]
    q2 = [cloned]
    while(q1.length > 0) {
        temp1 = []
        temp2 = []
        while(q1.length > 0) {
            n1 = q1.pop()
            n2 = q2.pop()
            if(n1.val === n2.val && n2.val === target.val) {
                return n2
            }
            n1.left && temp1.push(n1.left)
            n1.right && temp1.push(n1.right)
            n2.left && temp2.push(n2.left)
            n2.right && temp2.push(n2.right)
        }
        q1 = temp1
        q2 = temp2
    }
};
