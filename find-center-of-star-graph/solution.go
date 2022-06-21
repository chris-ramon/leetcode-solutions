// **Complexity Analysis**

// * Time: O(N), where N is the number of nodes.
// * Space: O(N), where N is the size of the adjacency list storing the nodes.

func findCenter(edges [][]int) int {
    adjacency := make(map[int][]int, 0)
    for i := 0; i < len(edges); i++ {
        n1, n2 := edges[i][0], edges[i][1]
        adjacency[n1] = append(adjacency[n1], n2)
        adjacency[n2] = append(adjacency[n2], n1)
    }
    max := 0
    result := 0
    for node, neighbors := range adjacency {
        if len(neighbors) > max {
            max = len(neighbors)
            result = node
        }
    }
    return result
}

/*
TestCases

[[1,2],[2,3],[4,2]]

[[1,2],[5,1],[1,3],[1,4]]
*/
