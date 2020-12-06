// **Complexity Analysis**
// - Time: `O(N)`, where N is the number of nodes in the tree.
// - Space: `O(N)`, because the size of the queue can grow up to N, where N is the number of nodes in the tree.
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
 * @param {number} k
 * @return {boolean}
 */
var findTarget = function (root, k) {
  if (root == null) return false;
  var queue = [root];
  var s = new Set();
  while (queue.length) {
    const size = queue.length;
    for (var i = 0; i < size; i++) {
      const node = queue.shift();
      const n = k - node.val;
      if (s.has(n)) {
        return true;
      }
      s.add(node.val);
      node.left && queue.push(node.left);
      node.right && queue.push(node.right);
    }
  }
  return false;
};

// Testcases:
/*
[]
0
[-5,1]
4
[5]
5
[5,3,6,2,4,null,7]
9
[5,3,6,2,4,null,7]
28
*/
