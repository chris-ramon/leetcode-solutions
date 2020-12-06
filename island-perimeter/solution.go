// Complexity Analysis
// Time: O(M*N) where M and N are the rows & columns from the grid.
// Space: O(M*N) where M x N is the size of the callstack during the dfs depth.

func islandPerimeter(grid [][]int) int {
    // 7 * 4 - (6 * 2) = 28 - 12 = 16
    // n * 4 - (n-1 * 2)
    // 4 * 4 - (3 * 2) = 16 - 6 = 10
    // 16 - (4-1 * 2) = 16 - 6 = 10
    if(len(grid) == 0) {
        return 0
    }
    
    var count int
    dfs(grid, 0, 0, &count)
    if count == 0 {
        return 0
    }
    // fmt.Printf("grid: %+v\n", grid)
    return (count * 4) - ((count - 1) * 2)
}

func dfs(grid [][]int, i int, j int, count *int) {
    // fmt.Printf("i: %v\n", i)
    // fmt.Printf("j: %v\n", j)
    // fmt.Printf("grid: %v\n", grid)
    if i < 0 ||j < 0 {
        return
    }
    if i > len(grid) - 1 || j > len(grid[i]) - 1 {
        return
    } 
    
    if grid[i][j] == -1 || (grid[i][j] == 0 && (i > 0 || j > 0) && *count > 0) {
        return
    }
    
    if grid[i][j] == 1 {
        *count += 1
    }
    
    grid[i][j] = -1
    
    dfs(grid, i - 1, j, count)
    dfs(grid, i, j + 1, count)
    dfs(grid, i + 1, j, count)
    dfs(grid, i, j - 1, count)
}

// Testcases:
/*
[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
[[0,1,0,1],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
[[0,1,0,1],[1,1,1,0],[0,1,0,0],[1,1,0,3]]
[[0,1,0,1],[1,1,1,0],[0,1,0,0],[1,1,0,-1]]
[]
[[0]]
[[0],[1],[1]]
[[0],[1],[1],[0,1]]
[[0],[1],[1],[0]]
[[-1],[-2]]
[[],[]]
[[0,0,1]]
[[0,0,0,0],[0,0,0,0],[0,0,0,1]]
[[1,0,0,0],[0,0,0,0],[0,0,0,1]]
[[0,0,0,0],[0,0,1,0],[0,0,0,0]]
[[1,1],[1,1]]
*/