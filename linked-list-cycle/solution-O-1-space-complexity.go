/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    var (
        fast *ListNode = head
        slow *ListNode = head
    )
    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if (fast == slow) {
            return true
        }
    }
    return false
}