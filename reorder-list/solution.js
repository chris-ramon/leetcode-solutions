/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {void} Do not return anything, modify head in-place instead.
 */
var reorderList = function(head) {
    if(!head || !head.next) return;
    
    var slow = head; var fast = head;
    
    while(fast != null && fast.next != null && fast.next.next != null) {
        fast = fast.next.next;
        slow = slow.next;
    }
    
    var secondHalf = slow.next;
    slow.next = null;
    
    merge(head, reverse(secondHalf));
};

var reverse = function(head) {
    var prev = null;
    while(head != null) {
        var temp = head.next;
        head.next = prev;
        prev = head;
        head = temp;
    }
    return prev;
};

var merge = function(l1, l2) {
    var l1Next = null; var l2Next = null;
    while(l2 != null) {
        l1Next = l1.next;
        l2Next = l2.next;
        
        l1.next = l2;
        l2.next = l1Next;
        
        l1 = l1Next;
        l2 = l2Next;
    }
};