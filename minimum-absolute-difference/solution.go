// **Complexity Analysis**
// * Time complexity: O(N * log N), where N is the length of the given slice.
// * Space complexity: O(N/2), where N is the size of the slice used to store the result.
import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Slice(arr, func(i int, j int) bool {
		return arr[j] > arr[i]
	})
	min := math.MaxFloat64
	keyPairs := map[float64]struct{}{
		float64(arr[0]): struct{}{},
	}
	for i := 1; i < len(arr); i++ {
		prev := arr[i-1]
		n := arr[i]
		diff := math.Abs(float64(n - prev))
		min = math.Min(min, diff)
		keyPairs[float64(n)] = struct{}{}
	}
	result := [][]int{}
	pairMap := map[string]struct{}{}
	for _, n := range arr {
		pair := float64(n) - min
		if n < 0 {
			pair = float64(n) + min
		}
		if pair == float64(n) {
			continue
		}
		if _, found := keyPairs[pair]; found {
			pairKey := fmt.Sprintf("%d:%d", n, int(pair))
			if _, found := pairMap[pairKey]; found {
				continue
			}
			if n < 0 {
				result = append(result, []int{n, int(pair)})
			} else {
				result = append(result, []int{int(pair), n})
			}
			pairMap[fmt.Sprintf("%d:%d", n, int(pair))] = struct{}{}
			pairMap[fmt.Sprintf("%d:%d", int(pair), n)] = struct{}{}
		}
	}
	return result
}
