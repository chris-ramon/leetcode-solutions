/**
 * // Definition for a Node.
 * function Node(val, next, random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */

/**
 * @param {Node} head
 * @return {Node}
 */
var copyRandomList = function(head) {
    var p = head;
    
    var sentinel = new Node();
    var temp = sentinel;
    
    var headNodes = new Map();
    var nodes = new Map();
    
    var i = -1;
    while(p != null) {
        i += 1;
        temp.next = new Node(p.val);
        
        headNodes.set(p, i);
        nodes.set(i, temp.next);
        
        temp = temp.next;
        p = p.next;
    }
    
    var p2 = head;
    var p3 = sentinel.next;
    
    var j = -1;
    while(p3 != null) {
        j += 1;
        
        if(p2.random != null) {
            var randomIdx = headNodes.get(p2.random);
            p3.random = nodes.get(randomIdx);
        } else {
            p3.random = null;
        }
        
        p2 = p2.next;
        p3 = p3.next;
    }
    
    return sentinel.next;
};