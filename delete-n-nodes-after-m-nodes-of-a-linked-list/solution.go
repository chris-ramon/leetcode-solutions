/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    var mCount int
    var nCount int
    var prev *ListNode
    var sentinel *ListNode = &ListNode{}
    sentinel.Next = head
    
    for head != nil {
        if prev == nil {
            mCount += 1
        } else {
            nCount += 1
        }
        
        if mCount == m {
            prev = head
            mCount = 0
        }
        
        if nCount == n {
            prev.Next = head.Next
            prev = nil
            nCount = 0
        }
        
        head = head.Next
    }
    
    if prev != nil && nCount > 0 {
        prev.Next = nil
    }
    
    return sentinel.Next
}