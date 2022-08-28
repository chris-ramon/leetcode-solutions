import "strconv"

func sortByBits(arr []int) []int {
	nums := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		n := strconv.FormatInt(int64(arr[i]), 2)
		count := 0
		for j := 0; j < len(n); j++ {
			if n[j] == '1' {
				count++
			}
		}
		nums[arr[i]] = count
	}
	sameCounts := true
	for i := 0; i < len(arr)-1; i++ {
		if nums[arr[i]] != nums[arr[i+1]] {
			sameCounts = false
			break
		}
	}
	if sameCounts {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	} else {
		sort.Slice(arr, func(i, j int) bool {
			if nums[arr[i]] == nums[arr[j]] {
				return arr[i] < arr[j]
			}
			return nums[arr[i]] < nums[arr[j]]
		})
	}
	return arr
}

/*
TestCases
[0,1,2,3,4,5,6,7,8]\n[1024,512,256,128,64,32,16,8,4,2,1]\n[7,0,1]\n[0,100]\n[0]\n[1111,7644,1107,6978,8742,1,7403,7694,9193,4401,377,8641,5311,624,3554,6631]
*/
