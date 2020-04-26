/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    var (
        temp *ListNode = head
        result float64
        exp float64
    )
    
    decimalValue(temp, &result, &exp)
    
    return int(result)
}

func decimalValue(node *ListNode, result *float64, exp *float64) {
    if node == nil {
        return
    }
    
    decimalValue(node.Next, result, exp)
    
    if node.Val == 1 {
        *result += math.Pow(2, *exp)
    }
    
    *exp += 1
}