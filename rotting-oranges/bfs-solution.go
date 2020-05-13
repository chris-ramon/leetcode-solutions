func orangesRotting(grid [][]int) int {
    if(len(grid) == 0) {
        return -1
    }
    
    var (
        queue [][]int = [][]int{}
        minutesElapsed int
        freshOranges int
        directions [][]int = [][]int{[]int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}}
    )
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            switch grid[i][j] {
                case 0:
                    continue
                case 1:
                    freshOranges += 1
                case 2:
                    queue = append(queue, []int{i, j})
            }
        }
    }
    
    if(freshOranges == 0) {
        return 0
    }
    
    
    for len(queue) > 0 {
        size := len(queue)
        
        // fmt.Printf("minutesElapsed: %v, grid: %+v \n", minutesElapsed, grid)
        
        for i := 0; i < size; i++ {
            cell := queue[0]
            queue = queue[1:]

            row, column := cell[0], cell[1]
            
            for j := 0; j < len(directions); j++ {
                d := directions[j]
                neighborRow, neighborCol := row + d[0], column + d[1]
                
                validNRow := neighborRow >= 0 && neighborRow < len(grid)
                validNColumn := validNRow && neighborCol >= 0 && neighborCol < len(grid[neighborRow])
                
                if validNRow && validNColumn && grid[neighborRow][neighborCol] == 1 {
                    grid[neighborRow][neighborCol] = 2
                    freshOranges -= 1
                    queue = append(queue, []int{neighborRow, neighborCol})
                    
                }
            }
        }
        
        if len(queue) > 0 {
            minutesElapsed += 1
        }
    }
    
    if(freshOranges > 0) {
        return -1
    }
    
    return minutesElapsed
}