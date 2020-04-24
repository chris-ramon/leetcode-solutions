/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    
    var (
        slow *ListNode = head
        fast *ListNode = head
    )
    
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next;
        slow = slow.Next;
    }
    
    slow.Next = reverse(slow.Next)
    slow = slow.Next
    
    for slow != nil {
        if slow.Val != head.Val {
            return false
        }
        slow = slow.Next
        head = head.Next
    }
    
    return true
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp
    }
    return prev
}