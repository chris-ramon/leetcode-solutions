/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    var s1 = l1; var s2 = l2;
    var q1 = []; var q2 = [];
    
    while(s1) {
        q1.push(s1.val);
        s1 = s1.next;
    }
    while(s2) {
        q2.push(s2.val);
        s2 = s2.next;
    }
    
    var sentinel = new ListNode();
    var temp = sentinel;
    var carry = false;
    while(q1.length > 0 || q2.length > 0) {
        var n1 = q1.length > 0 ? q1.pop(): 0;
        var n2 = q2.length > 0 ? q2.pop(): 0;
        var sum = n1 + n2;
        
        sum = carry ? sum + 1 : sum;
        carry = sum >= 10;
        
        var n = new ListNode(sum % 10);
        temp.next = n;
        temp = temp.next;
    }
    
    if(carry) {
        temp.next = new ListNode(1);
    }
    
    var head = sentinel.next;
    var prev = null;
    while(head) {
        var t = head.next;
        head.next = prev;
        prev = head;
        head = t;
    }
    
    return prev;
};