/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var (
        prev *ListNode 
        sentinel *ListNode = &ListNode{Next: head}
        idx int
    )
    
    for head != nil {
        idx += 1
        
        if head.Next != nil && idx + 1 == m {
            prev = head
        }
        
        if idx != m {
            head = head.Next
            continue
        }
        
        newHead := reverse(head, n - m)
        
        if prev != nil {
            prev.Next = newHead
        } else {
            sentinel.Next = newHead
        }
        
        head = nil
    }
    
    return sentinel.Next
}

func reverse(head *ListNode, i int) *ListNode {
    var (
        initialHead *ListNode = head
        prev *ListNode
        idx int
    )
    
    for head != nil {
        idx += 1
        
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp 
        
        if idx > i {
            initialHead.Next = head
            head = nil
        }
    }
    return prev
}