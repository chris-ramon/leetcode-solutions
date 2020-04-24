/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var (
        sentinel *ListNode = &ListNode{}
        prev *ListNode = sentinel
        p *ListNode = head
    )
    
    sentinel.Next = head
    
    for p != nil {
        if p.Val == val {
            prev.Next = p.Next
        } else {
            prev = p
        }
        p = p.Next
    }
    
    return sentinel.Next
}