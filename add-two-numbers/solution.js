/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    var result = new ListNode();
    if (!l1 && !l2) return result;
    var p1 = l1; var p2 = l2; var p3 = result; var carry = false;
    while(p1 != null || p2 != null) {
        var val1 = (p1 && p1 != null) ? p1.val : 0;
        var val2 = (p3 && p2 != null) ? p2.val : 0;
        var val3 = carry ? 1 : 0;
        carry = false;
        
        var sum = val1 + val2 + val3;
        if (sum >= 10) {
            sum = sum % 10;
            carry = true;
        }
        
        p3.next = new ListNode(sum);
        p3 = p3.next;
        
        p1 = (p1 && p1.next != null) ? p1.next : null;
        p2 = (p2 && p2.next != null) ? p2.next : null;
    }
    if (carry) p3.next = new ListNode(1);
    return result.next;
};
