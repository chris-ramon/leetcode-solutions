/*
Complexity Analysis:
Time Complexity: O(N * logN), where N is the length of the arr1.
Space Complexity: O(M + N), where M + N is the size of the arr1 & arr2.
*/
package main

import (
    "slices"
    "sort"
    "fmt"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
    sort.Ints(arr1)
    // fmt.Printf("arr1: %+v\n", arr1)
    // fmt.Printf("arr2: %+v\n", arr2)
    arr1Map := map[int]int{}
    for i := 0; i < len(arr1); i++ {
        if _, found := arr1Map[arr1[i]]; found {
            arr1Map[arr1[i]]++
        } else {
            arr1Map[arr1[i]] = 1
        }
    }
    // fmt.Printf("arr1Map: %+v\n", arr1Map)
    arr2Map := map[int]int{}
    for i := 0; i < len(arr2); i++ {
        if _, found := arr2Map[arr2[i]]; !found {
          arr2Map[arr2[i]]++
        } else {
          arr2Map[arr2[i]] = 1
        }
    }
    notFound := []int{}
    result := []int{}
    for i := 0; i < len(arr1); i++ {
        if _, found := arr2Map[arr1[i]]; !found {
            notFound = append(notFound, arr1[i])
        }
    }
    // fmt.Printf("notFound: %+v\n", notFound)
    // fmt.Printf("arr2Map: %+v\n", arr2Map)
    for i := 0; i < len(arr2); i++ {
        count, _ := arr1Map[arr2[i]]
        for j := 0; j < count; j++ {
            result = append(result, arr2[i])
        }
    }
    result = slices.Concat(result, notFound)
    // fmt.Printf("result: %+v\n", result)
    // fmt.Printf("notFound: %+v\n", notFound)
    return result
}
