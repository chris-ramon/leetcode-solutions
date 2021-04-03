/**
Complexity Analysis
- Time: O(N)
- Space: O(N)
*/


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
    oldAndNewMap := make(map[*Node]*Node)
    
    sentinel := Node{Next: &Node{}}
    temp := sentinel.Next
    
    for head != nil {
        temp.Val = head.Val
        temp.Random = head.Random
        
        if head.Next != nil {
            temp.Next = &Node{}
        }
        
        oldAndNewMap[head] = temp
        
        temp = temp.Next
        head = head.Next
    }
    
    temp2 := sentinel.Next
    
    for temp2 != nil {
        newRandom, _ := oldAndNewMap[temp2.Random]
        temp2.Random = newRandom
        temp2 = temp2.Next
    }
    
    return sentinel.Next
}
