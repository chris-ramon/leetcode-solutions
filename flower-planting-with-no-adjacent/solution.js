func gardenNoAdj(N int, paths [][]int) []int {
    var edges map[int]map[int]bool = make(map[int]map[int]bool, 0)
    
    for i := 0; i < len(paths); i++ {
        from, to := paths[i][0], paths[i][1]
        if _, found := edges[from]; found {
            nodes := edges[from]
            nodes[to] = true
            edges[from] = nodes
        } else {
            nodes := make(map[int]bool, 1)
            nodes[to] = true
            edges[from] = nodes
        }
        if _, found := edges[to]; found {
            nodes := edges[to]
            nodes[from] = true
            edges[to] = nodes
        } else {
            nodes := make(map[int]bool, 1)
            nodes[from] = true
            edges[to] = nodes
        }
    }
    
    var totalFlowers int = 4
    var result []int
    var queue []int
    var flowerByGarden map[int]int = make(map[int]int, 0)
    
    for i := 1; i <= N; i++ {
        queue = append(queue, i)
        flowerByGarden[i] = 0
    }
    
    for len(queue) > 0 {
        garden := queue[0]
        queue = queue[1:]
        gardenPaths := edges[garden]
        
        OuterLoop:
        for f := 1; f <= totalFlowers; f++ {
            for connectedGarden, _ := range gardenPaths {
                if flowerByGarden[connectedGarden] == f {
                    continue OuterLoop
                }
            }
            flowerByGarden[garden] = f
            result = append(result, f)
            break
        }
    }
    
    return result
}