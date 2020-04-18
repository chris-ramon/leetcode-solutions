/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
    var result = 0;
    if (!grid || !grid.length) return result;
    for (var i = 0; i < grid.length; i++) {
        for (var j = 0; j < grid[i].length; j++) {
            if (grid[i][j] !== "1") continue;
            result += 1;
            dfs(grid, i, j);
        }
    }
    return result;
};

const dfs = (grid, i, j) => {
    if (!grid || i < 0 || j < 0 || i > grid.length - 1 || grid[i] === undefined || j > grid[i].length - 1 || grid[i][j] === "0") {
        return;
    }
    grid[i][j] = "0";
    dfs(grid, i + 1, j);
    dfs(grid, i, j + 1);
    dfs(grid, i - 1, j);
    dfs(grid, i, j - 1);
};
