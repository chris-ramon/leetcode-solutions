/*
Complexity Analysis:
- Time: O(n), n size of the given array.
- Space: O(1), constant number of variables used.
*/

func threeConsecutiveOdds(arr []int) bool {
    count := 0
    for i := 0; i < len(arr); i++ {
        if arr[i] % 2 != 0 {
            count++
        } else {
            count = 1
        }
        if count == 3 {
            return true
        }
    }
    return false
}
