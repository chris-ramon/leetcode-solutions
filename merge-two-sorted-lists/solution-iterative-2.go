/*
Complexity Analysis
- Time: O(N+M) where N and M are the length of the given linked lists.
- Space: O(N+M) N+M is the size of the returned new linked list.
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var (
        p1 *ListNode = l1
        p2 *ListNode = l2
        p3 *ListNode = &ListNode{}
        sentinel *ListNode
    )
    
    sentinel = p3
    
    for p1 != nil || p2 != nil {
        var v1 *int
        var v2 *int
        
        if p1 != nil {
            v1 = &p1.Val
        }
        
        if p2 != nil {
            v2 = &p2.Val
        }
        
        if v1 != nil && v2 != nil && *v1 <= *v2 {
            p3.Next = &ListNode{Val: *v1}
            p3 = p3.Next
            p1 = p1.Next
            continue
        }
        
        if v1 != nil && v2 != nil && *v1 > *v2 {
            p3.Next = &ListNode{Val: *v2}
            p3 = p3.Next
            p2 = p2.Next
            continue
        }
        
        if v1 != nil {
            p3.Next = &ListNode{Val: *v1}
            p3 = p3.Next           
            p1 = p1.Next
        }
        
        if v2 != nil {
            p3.Next = &ListNode{Val: *v2}
            p3 = p3.Next           
            p2 = p2.Next
        }
    }
    
    // fmt.Printf("p1: %+v\n", p1)
    // fmt.Printf("p2: %+v\n", p2)
    
    return sentinel.Next
}

/*
Testcases:
[1,2,4]
[1,3,4]

[]
[]

[-1]
[]

[0]
[-1]

[0]
[0]

[1,2,3]
[]

[]
[3,33,333]

[1,2,3,4]
[1,2,3,4,5]

[1,2,3,4]
[10]
*/