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
 * @return {number[][]}
 */
var verticalOrder = function(root) {
    var valsByColumn = new Map();
    var columns = [];
    var result = [];
    
    var queue = [{node: root, column: 0}];
    while(queue.length) {
        var {node, column} = queue.shift();
        
        if(node === null) continue;
        
        if(valsByColumn.has(column)) {
            var vals = valsByColumn.get(column);
            vals.push(node.val);
            valsByColumn.set(column, vals);
        } else {
            valsByColumn.set(column, [node.val]);
        }
        
        queue.push({node: node.left, column: column - 1});
        queue.push({node: node.right, column: column + 1});
    }
    
    for(let [k, v] of valsByColumn) {
        columns.push(k);
    }
    
    columns.sort((a,b) => a - b);
    
    for(var i = 0; i < columns.length; i++) {
        var column = columns[i];
        var vals = valsByColumn.get(column);
        result.push(vals);
    }
    
    return result;
};