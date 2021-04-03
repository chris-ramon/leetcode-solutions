/*
Complexity Analysis:
- Time: O(N), where N are the length of the given nums.
- Space: O(N), where N are the length of the given nums, that are stored in the hash-map/table.
*/

func twoSum(nums []int, target int) []int {
    cache := make(map[int]int, len(nums))
    for i, n := range nums {
        diff := target - n
        idx, found := cache[diff]
        if found {
            return []int{idx, i}
        }
        cache[n] = i
    }
    return []int{}
}
