// Complexity Analysis
// Time: O(N), where N is the number of nums.
// Space: O(1)

func singleNumber(nums []int) int {
    // 2 ^ 2
    // 10 ^
    // 10
    // 00
    
    // 0 ^ 1
    // 00 ^
    // 01
    // 01 -> 1

    // 4 ^ 1
    // 100 ^
    // 001
    // 101
    
    // 5 ^ 2
    // 101 ^
    // 010
    // 111
    
    // 7 ^ 1
    // 111 ^
    // 001
    // 110
    
    // 6 ^ 2
    // 110 ^
    // 010
    // 100 -> 4
    
    // Applying XOR of duplicated numbers become 0.
    // At the end XOR keeps the non unique number.
    
    var result int
    for _, n := range nums {
        result ^= n
    }
    return result
}

// Testcases:
/*
[2,2,1]
[4,1,2,1,2]
[]
[1]
*/
