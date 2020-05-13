/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */
var minKnightMoves = function(x, y) {
    if(x == 0 && y === 0) return 0;
    var queue = [[0, 0]];
    var knightMoves = [[-2,1],[-1,2],[1,2],[2,1],[2,-1],[1,-2],[-1,-2],[-2,-1]];
    var visited = new Set();
    var memo = new Map();
    memo.set("0-0", 0);
    var totalMoves = 0;
    while(queue.length) {
        var size = queue.length;
        totalMoves += 1;
        for(var i = 0; i < size; i++) {
            var [px, py] = queue.shift();
            for(var j = 0; j < knightMoves.length; j++) {
                var [kx, ky] = knightMoves[j];
                var nx = px + kx;
                var ny = py + ky;
                var next = nx + "-" + ny;
                if(!memo.has(next)) {
                    memo.set(next, totalMoves);
                    queue.push([nx, ny]);
                    if(nx === x && ny === y) {
                        return memo.get(x + "-" + y);
                    }
                }
            }
        }
    }
    return memo.get(x + "-" + y);
};