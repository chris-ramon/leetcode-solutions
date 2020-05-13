func minKnightMoves(x int, y int) int {
    if(x == 0 && y == 0) {
        return 0
    }
    
    var(
        totalMoves int
        /*
        directions [][]int = [][]int{
            []int{-1,2},
            []int{-1,-2},
            []int{-2,1},
            []int{-2,-1},
            []int{1,2},
            []int{1,-2},
            []int{2,1},
            []int{2,1},
        }
        */
        directions [][]int = [][]int{
            []int{2,1},
            []int{1,2},
            []int{-1,2},
            []int{-2,1},
            []int{-2,-1},
            []int{-1,-2},
            []int{1,-2},
            []int{2,-1},
        }
        visited map[string]bool = make(map[string]bool, 0)
        queue [][]int = [][]int{[]int{0, 0}}
        visistedKey func(x, y int) string
    )
    
    visistedKey = func(x, y int) string {
        return fmt.Sprintf("%d%d", x, y)
    }
    
    for len(queue) > 0 {
        size := len(queue)
        
        totalMoves += 1
        
        for i := 0; i < size; i++ {
            coords := queue[0]
            queue = queue[1:]
            
            cx, cy := coords[0], coords[1]
            
            for j := 0; j < len(directions); j++ {
                d := directions[j]
                dx, dy := d[0], d[1]
                nx, ny := cx + dx, cy + dy
                
                fmt.Printf("visited: %+v \n", visited)
                
                if nx == x && ny == y {
                    return totalMoves
                }
                
                if _, found := visited[visistedKey(nx, ny)]; !found {
                    visited[visistedKey(nx, ny)] = true
                    queue = append(queue, []int{nx, ny})
                }
                
                fmt.Printf("queue: %+v \n", queue)
            }
        }
    }
    
    return 0
}