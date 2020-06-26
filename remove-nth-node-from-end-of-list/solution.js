/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    var p1 = head; var p2 = head; var len = 0;
    while(p1) {
        len += 1;
        p1 = p1.next;
    }
    if(n > len) n = len;
    var i = 1;
    var prev = null;
    while(p2) {
        if(i === len - n + 1) {
            if(prev) prev.next = p2.next;
            if(prev === null) head = p2.next;
            break;
        }
        i += 1;
        prev = p2;
        p2 = p2.next;
    }
    return head;
};