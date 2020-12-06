// **Complexity Analysis**
// - Time: O(N) where N is the number of nodes in the tree.
// - Space: O(N) where N is the number of nodes in the tree.

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} target
 * @param {number} K
 * @return {number[]}
 */
var distanceK = function (root, target, K) {
  var result = [];

  var queue = [root];
  var targetNode;

  while (queue.length) {
    var size = queue.length;

    for (var i = 0; i < size; i++) {
      const node = queue.shift();

      if (target && target.val == node.val) {
        targetNode = node;
      }

      if (node.left) {
        node.left.parent = node;
        queue.push(node.left);
      }

      if (node.right) {
        node.right.parent = node;
        queue.push(node.right);
      }
    }
  }

  if (!targetNode) return result;

  var queueK = [targetNode];
  var visited = new Set();
  var level = 0;

  while (queueK.length) {
    var size = queueK.length;

    for (var i = 0; i < size; i++) {
      const node = queueK.shift();

      if (visited.has(node.val)) continue;

      node.parent && !visited.has(node.parent.val) && queueK.push(node.parent);
      node.left && !visited.has(node.left.val) && queueK.push(node.left);
      node.right && !visited.has(node.right.val) && queueK.push(node.right);

      visited.add(node.val);

      if (level == K) {
        result.push(node.val);
      }
    }

    level += 1;
  }

  return result;
};

// Test Cases
/*
[3,5,1,6,2,0,8,null,null,7,4]
0
5

[3,5,1,6,2,0,8,null,null,7,4]
5
2

[3,5,1,6,2,0,8,null,null,7,4]
3
0

[3,5,1,6,2,0,8,null,null,7,4]
0
2
*/
