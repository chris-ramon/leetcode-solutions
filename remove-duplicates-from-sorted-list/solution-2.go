/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var current *ListNode = head
    for current != nil && current.Next != nil {
        if(current.Val == current.Next.Val) {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return head
}
