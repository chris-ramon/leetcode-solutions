/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function(head) {
    if (head == null || head.next == null) return true;
    
    var slow = head;
    var fast = head;
    
    while(fast.next != null && fast.next.next != null) {
        fast = fast.next.next;
        slow = slow.next;
    }
    
    slow.next = reverseLinkedList(slow.next);
    slow = slow.next;
    
    while (slow != null) {
        if (head.val != slow.val) return false;
        head = head.next;
        slow = slow.next;
    }
    
    return true;
};

const reverseLinkedList = (l) => {
    var prev = null;
    while (l != null) {
        var temp = l.next;
        l.next = prev;
        prev = l;
        l = temp;
    }
    return prev;
};