// Complexty Analysis
// Time: `O(N)`, where N is the number of iterations to build dp.
// Space: `O(N)`, where N is the size of the dp map.

func fib(N int) int {
    var dp map[int]int = make(map[int]int, N + 1)
    
    dp[0] = 0
    dp[1] = 1
    
    for i := 2; i <= N; i++ {
        dp[i] = dp[i - 2] + dp[i - 1]
    }
    
    return dp[N]
}

// Testcases:
/*
0
1
2
3
4
11
*/

// Complexty Analysis
// Time: `O(2^N)`, where N is the given number which doubles on each level of recursion.
// Space: `O(N)`, where N is the size of the callstack during the recursion.

// func fib(N int) int {
//     if N == 0 {
//         return 0
//     }
//     if N == 1 {
//         return 1
//     }
//     return fib(N - 1) + fib(N -2)
// }

// Testcases:
/*
0
1
2
3
4
11
*/