// **Complexity Analysis**

// * Time Complexity: O(N) where N is the number of nodes in the graph.
// * Space Complexity: O(N) where N is the que size during the breadth first search.

func validPath(n int, edges [][]int, source int, destination int) bool {
    adjacency := make(map[int][]int, 0)
    for i := 0; i < len(edges); i++ {
        n1, n2 := edges[i][0], edges[i][1]
        adjacency[n1] = append(adjacency[n1], n2)
        adjacency[n2] = append(adjacency[n2], n1)
    }
    queue := []int{source}
    visited := make(map[int]bool, 0)
    for len(queue) > 0 {
        temp := []int{}
        for len(queue) > 0 {
            n := queue[0]
            if n == destination {
                return true
            }
            queue = queue[1:]
            if _, found := visited[n]; found {
                continue
            }
            neighbors := adjacency[n]
            temp = append(temp, neighbors...)
            visited[n] = true
        }
        queue = temp
    }
    return false
}
