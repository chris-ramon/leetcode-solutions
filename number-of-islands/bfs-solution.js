/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
    var result = 0;
    if (!grid || !grid.length) return result;
    for (var i = 0; i < grid.length; i++) {
        for (var j = 0; j < grid[i].length; j++) {
            if (grid[i][j] === "0") continue;
            result += 1;
            bfs(grid, i, j);
        }
    }
    return result;
};

var bfs = function(grid, i, j) {
	var queue = [{x: i, y: j}];
	while(queue.length) {
	  var size = queue.length;
	  for (var l = 0; l < size; l++) {
	    var land = queue.shift();
			var { x, y } = land;
	  	  if (x < 0 || y < 0 || x >= grid.length || y >= grid[x].length || grid[x][y] === "0") {
				continue;
			}
			grid[x][y] = "0";
			var top = {x: x + 1, y};
			var right = {x, y: y + 1};
			var bottom = {x: x - 1, y};
			var left = {x, y: y - 1};
			queue.push(top, right, bottom, left);
	  }
	}
}
