// Complexity Analysis
// Time: `O(N)`, where N is the size of the given input.
// Space: `O(N)`, where N is the size of the dp map.

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    var dp map[int]float64 = make(map[int]float64, len(nums) - 1)
    dp[0] = float64(nums[0])
    dp[1] = math.Max(float64(nums[0]), float64(nums[1]))
    
    for i := 2; i < len(nums); i++ {
        dp[i] = math.Max(float64(nums[i]) + dp[i - 2], dp[i - 1])
    }

    return int(dp[len(nums) - 1])
}

// Testcases:
/*
[]
[1]
[1,5]
[1,2,3,1]
[2,7,9,3,1]
[1,2,3,4,5,6,7,8,9,10,11,100,300]
*/