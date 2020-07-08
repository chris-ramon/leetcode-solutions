/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		slow     *ListNode = head
		slowPrev *ListNode = head
		fast     *ListNode = head
	)

	for fast != nil && fast.Next != nil {
		slowPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowPrev.Next = nil

	leftList := sortList(head)
	rightList := sortList(slow)

	return merge(leftList, rightList)
}

func merge(l, r *ListNode) *ListNode {
	var (
		p        *ListNode = &ListNode{}
		sentinel *ListNode = &ListNode{Next: p}
	)

	for l != nil && r != nil {
		if l.Val <= r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}

	if l != nil {
		p.Next = l
	}
	if r != nil {
		p.Next = r
	}

	return sentinel.Next.Next
}
