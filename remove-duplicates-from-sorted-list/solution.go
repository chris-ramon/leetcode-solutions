/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var (
        current *ListNode = head
        prev *ListNode
    )
    for current != nil {
        if (prev != nil && current.Val == prev.Val) {
            prev.Next = current.Next
        } else {
            prev = current
        }
        current = current.Next
    }
    return head
}