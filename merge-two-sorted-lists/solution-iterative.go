/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode = &ListNode{}
    var head *ListNode = result
    
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            head.Next = l1
            l1 = l1.Next
        } else {
            head.Next = l2
            l2 = l2.Next
        }
        head = head.Next
    }
    
    if l1 != nil {
        head.Next = l1
    } else {
        head.Next = l2
    }
    
    return result.Next
}