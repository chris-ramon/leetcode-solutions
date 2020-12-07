// Complexity Analysis
// - Time: O(N ^ H/2), where N is the number of square numbers and H is the height of the tree.
// - Space: O(sqrt(N) ^ H), where N is the number of square numbers and H is the height of the tree.

/**
 * @param {number} n
 * @return {number}
 */
var numSquares = function(n) {
    if(n == 0) return 0
    
    var squares = []
    var queue = [n]
    var level = 1
    
    for(var i = 1; i <= n; i++) {
        const num = i ** 2
        if(num > n) break
        squares.push(num)
    }
    
    while(queue.length) {
        var size = queue.length
        var visited = new Set()
        
        for(var i = 0; i < size; i++) {
            var num = queue.shift()
            
            for(var sq of squares) {
                
                const diff = num - sq
                
                if(diff < 0) continue
                
                if(diff === 0) {
                    return level
                }
                
                if(!visited.has(diff)) {
                    visited.add(diff)
                    queue.push(diff)
                }
            }
        }
        
        level += 1
    }
    
    return level
    
};

// Testcases:
/*
0
1
2
3
12
13
100
*/


// n = 12
// 1, 4, 9

// 10, 8, 3 - 1

// 9, 7, 2 - 2-1
// 6, 4, -1 - 2-2
// 1, -1, -6 - 2-3

// 8, 6, 2 - 3-1
// 5, 3, -1 - 3-2
// 0, -2, -6 - 3-3

// 5, 3, 2 - 3-4
// 2, *0, -1 - 3-5
// -3, -5, -6 - 3-6

// 0, 7, 2 - 3-7
// -3, 4, -1 - 3-8
// -8, -1, -6 - 3-9