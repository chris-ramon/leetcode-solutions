/*
Complexity Analysis
- Time:
  - Add: O(1) constant number of variables used.
  - Find: O(n) where n is the number of numbers added.
- Space:
  - Add: O(1).
  - Find: O(n) where n is the nums size.
*/


type TwoSum struct {
	sums map[int]int
	nums map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		sums: make(map[int]int, 0),
		nums: make(map[int]int, 0),
	}
}

func (this *TwoSum) Add(number int) {
	this.nums[number]++
}

func (this *TwoSum) Find(value int) bool {
	for num, count := range this.nums {
		complement := value - num
		if complement != num {
			if _, found := this.nums[complement]; found {
				return true
			}
		} else {
			if count > 1 {
				return true
			}
		}
	}
	return false
}

/*
TestCases
["TwoSum", "add", "add", "add", "find", "find"]
[[], [1], [3], [5], [4], [7]]
["TwoSum", "add", "add", "add", "find", "find"]
[[], [100], [300], [500], [4], [7]]
["TwoSum","add","find"]
[[],[0],[0]]
["TwoSum","add","add","find"]
[[],[0],[0],[0]]
*/

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
