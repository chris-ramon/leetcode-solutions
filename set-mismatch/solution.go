// # Complexity
// - Time complexity: `O(N)`, where N is the size of given nums. 
// - Space complexity: `O(N)`, where N is the size of given nums. 

func findErrorNums(nums []int) []int {
	numsMap := map[int]int{}
	for _, num := range nums {
		numsMap[num]++
	}
	missing := -1
	duplicated := -1
	for i := 1; i <= len(nums); i++ {
		if count, found := numsMap[i]; found {
			if count >= 2 {
				duplicated = i
			}
		} else {
			missing = i
		}
	}
	return []int{duplicated, missing}
}
