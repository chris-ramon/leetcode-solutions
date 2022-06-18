// **Complexity Analysis**
// - Time Complexity: O(N) where n is each cell within the grid.
// - Space Complexity: O(N) where n is the queue size during the breadth first search.

func islandPerimeter(grid [][]int) int {
    result := 0
    q := [][]int{}
    for i := 0; i < len(grid); i++ {
        if i == 1 {
            break
        }
        for j := 0; j < len(grid[i]); j++ {
            q = append(q, []int{i, j})
        }
    }
    for len(q) > 0 {
        temp := [][]int{}
        for len(q) > 0 {
            n := q[0]
            q = q[1:]
            x, y := n[0], n[1]
            count := countPerimeter(grid, x, y)
            result += count
            next := nextCell(grid, x, y)
            if len(next) > 0 {
                temp = append(temp, next)
            }
        }
        q = temp
    }
    return result
}

func countPerimeter(grid [][]int, x int, y int) int {
    if grid[x][y] == 0 {
        return 0
    }
    result := 0
    // up
    if (x-1 >= 0 && grid[x-1][y] == 0) || x-1 < 0 {
        result += 1
    }
    // down
    if (x+1 <= len(grid)-1 && grid[x+1][y] == 0) || x+1 > len(grid)-1 {
        result += 1
    }
    // left
    if (y-1 >= 0 && grid[x][y-1] == 0) || y-1 < 0 {
        result += 1
    }
    // right
    if (y+1 <= len(grid[x])-1 && grid[x][y+1] == 0) || y+1 > len(grid[x])-1 {
        result += 1
    }
    return result
}

func nextCell(grid [][]int, x int, y int) []int {
    if x >= len(grid)-1 {
        return []int{}
    }
    return []int{x+1,y}
}
