/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result ListNode = ListNode{}
    if l1 == nil && l2 == nil {
        return result.Next
    }
    
    var p1 *ListNode = l1
    var p2 *ListNode = l2
    var p3 *ListNode = &result
    var carry bool = false
    
    for {
        if p1 == nil && p2 == nil {
            break
        }
        
        var sum int = 0
        if p1 != nil {
            sum += p1.Val
        }
        if p2 != nil {
            sum += p2.Val
        }
        if carry {
            sum += 1
            carry = false
        }
        
        if sum >= 10 {
            sum = sum % 10
            carry = true
        }
        
        p3.Next = &ListNode{Val: sum}
        p3 = p3.Next
        
        if p1 != nil && p1.Next != nil {
            p1 = p1.Next
        } else {
            p1 = nil
        }
        
        if p2 != nil && p2.Next != nil {
            p2 = p2.Next
        } else {
            p2 = nil
        }
    }
    
    if carry {
        p3.Next = &ListNode{Val: 1} 
    }    
    
    return result.Next
}