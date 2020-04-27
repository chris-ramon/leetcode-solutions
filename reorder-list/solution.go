/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || (head != nil && head.Next == nil) {
		return
	}

	var (
		slow               *ListNode = head
		fast               *ListNode = head
		firstHalf          *ListNode = head
		secondHalfReversed *ListNode
	)

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalfReversed = reverse(slow.Next)

	slow.Next = nil

	merge(firstHalf, secondHalfReversed)
}

func reverse(ln *ListNode) *ListNode {
	var prev *ListNode
	for ln != nil {
		temp := ln.Next
		ln.Next = prev
		prev = ln
		ln = temp
	}
	return prev
}

func merge(ln1 *ListNode, ln2 *ListNode) {
	for ln1 != nil && ln2 != nil {
		ln1Next := ln1.Next
		ln2Next := ln2.Next

		ln1.Next = ln2
		ln2.Next = ln1Next

		ln1 = ln1Next
		ln2 = ln2Next
	}
}