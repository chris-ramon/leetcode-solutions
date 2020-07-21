// Complexity Analysis
// Time: `O(N)` where N is the number of rows within given costs.
// Space: `O(1)` in-place solution.
func minCost(costs [][]int) int {
    if len(costs) == 0 {
        return 0
    }
    for i := 1; i < len(costs); i++ {
        costs[i][0] += int(math.Min(float64(costs[i - 1][1]), float64(costs[i - 1][2])))
        costs[i][1] += int(math.Min(float64(costs[i - 1][0]), float64(costs[i - 1][2])))
        costs[i][2] += int(math.Min(float64(costs[i - 1][0]), float64(costs[i - 1][1])))
    }
    
    return int(math.Min(math.Min(float64(costs[len(costs)-1][0]), float64(costs[len(costs)-1][1])), float64(costs[len(costs)-1][2]))) 
}


// Testcases:
/*
[]
[[17,2,17],[16,16,5],[14,3,19]]
[[1,2,3],[0,0,0],[0,0,0]]
[[3,2,1],[0,0,0],[0,0,0]]
*/