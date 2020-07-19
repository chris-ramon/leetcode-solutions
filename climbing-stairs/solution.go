// Complexity Analysis
// Time: `O(N)`, where N is the number of iterations to build the dp map.
// Space: `O(N)`, where N is the size of the dp map.

func climbStairs(n int) int {
    var dp map[int]int = make(map[int]int, n + 1)
    
    dp[0] = 1
    dp[1] = 1
    
    for i := 2; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    
    return dp[n]
}

// Testcases:
/*
1
2
3
4
10
11
*/
