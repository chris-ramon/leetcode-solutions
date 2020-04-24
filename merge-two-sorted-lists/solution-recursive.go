/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var p3 *ListNode = &ListNode{}
    
    var p1 *ListNode = l1
    var p2 *ListNode = l2
    
    merger(p1, p2, p3)
    
    return p3.Next
}

func merger(p1 *ListNode, p2 *ListNode, p3 *ListNode) {
    if p1 == nil && p2 == nil {
        return
    }
    
    if p1 != nil && p2 != nil && p1.Val <= p2.Val {
        p3.Next = &ListNode{Val: p1.Val}
        p3 = p3.Next
        
        p1 = p1.Next
        
        merger(p1, p2, p3)
    } else if p1 != nil && p2 != nil && p1.Val >= p2.Val {
        p3.Next = &ListNode{Val: p2.Val}
        p3 = p3.Next
        
        p2 = p2.Next
        
        merger(p1, p2, p3)
    } else if p1 != nil && p2 == nil {
        p3.Next = &ListNode{Val: p1.Val}
        p3 = p3.Next
        
        p1 = p1.Next
        
        merger(p1, p2, p3)
    } else if p2 != nil && p1 == nil {
        p3.Next = &ListNode{Val: p2.Val}
        p3 = p3.Next
        
        p2 = p2.Next
        
        merger(p1, p2, p3)
    }
}