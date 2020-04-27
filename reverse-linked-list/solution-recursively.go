/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    return reverse(head)
}

func reverse(n *ListNode) *ListNode {
    if n == nil || n.Next == nil {
        return n
    }

    p := reverse(n.Next)
    
    n.Next.Next = n
    n.Next = nil
    
    return p
}
