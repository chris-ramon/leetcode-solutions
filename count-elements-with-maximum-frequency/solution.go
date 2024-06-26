package main

import "fmt"

func main() {
  fmt.Println(maxFrequencyElements([]int{10,12,11,9,6,19,11}))
}

func maxFrequencyElements(nums []int) int {
    numsMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if _, found := numsMap[nums[i]]; !found {
            numsMap[nums[i]] = 1
        }
    }
    return len(numsMap)
}
