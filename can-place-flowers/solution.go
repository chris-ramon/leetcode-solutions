// **Complexity Analysis**
// * Time: O(N), where N is the size of the given array.
// * Space: O(1), constant number of variables used.

func canPlaceFlowers(flowerbed []int, n int) bool {
	isLeftEmpty := func(i int) bool {
		if i == 0 {
			return true
		}
		return flowerbed[i-1] == 0
	}
	isRightEmpty := func(i int) bool {
		if i == len(flowerbed)-1 {
			return true
		}
		return flowerbed[i+1] == 0
	}
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && isLeftEmpty(i) && isRightEmpty(i) {
			flowerbed[i] = 1
			count += 1
            if count >= n {
                return true
            }
		}
	}
	return count >= n
}
