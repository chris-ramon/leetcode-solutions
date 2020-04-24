/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    if (headA == null || headB == null) return null;
    var p1 = headA; var p2 = headB; var visited = new Set();
    while(p1 != null) {
        visited.add(p1);
        p1 = p1.next;
    }
    while(p2 != null) {
        if(visited.has(p2)) {
            return p2;
        }
        p2 = p2.next;
    }
    return null;
};