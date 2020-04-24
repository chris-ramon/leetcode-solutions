/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    var sentinel = new ListNode();
    sentinel.next = head;
    
    var prev = sentinel;
    var current = head;
    
    while(current != null) {
        if(current.val == val) {
            prev.next = current.next;
        } else {
            prev = current;
        }
        current = current.next;
    }
    
    return sentinel.next;
};