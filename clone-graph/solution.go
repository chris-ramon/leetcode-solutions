/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    var (
        head *Node
        queue []*Node = []*Node{node}
        nodes map[*Node]*Node = make(map[*Node]*Node, 0)
        visited map[*Node]bool = make(map[*Node]bool, 0)
    )
    
    for len(queue) > 0 {
        n := queue[0]
        queue = queue[1:]
        
        if n == nil {
            continue
        }
        
        if _, found := visited[n]; found {
            continue
        }
        
        nCopy, found := nodes[n]
        if !found {
            nCopy = &Node{Val: n.Val, Neighbors: []*Node{}} 
        }
        
        visited[n] = true
        
        if head == nil {
            head = nCopy
        }
        
        nodes[n] = nCopy
        
        for i := 0; i < len(n.Neighbors); i++ {
            neighbor := n.Neighbors[i]
            neighborCopy, found := nodes[neighbor]
            if !found {
                newNeighborCopy := &Node{Val: neighbor.Val, Neighbors: []*Node{}} 
                nodes[neighbor] = newNeighborCopy
                neighborCopy = newNeighborCopy
            }
            nCopy.Neighbors = append(nCopy.Neighbors, neighborCopy)
        }
        
        queue = append(queue, n.Neighbors...)
    }
    
    return head
}