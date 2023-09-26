/*
Complexity Analysis:
- Time: O(1), constant runtime.
- Space: O(1), constant number of variables used.
*/


/**
 * @param {number} n
 * @return {Function} counter
 */
var createCounter = function(n) {
    let m = n - 1;
    return function() {
        return ++m;
    };
};

/** 
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */
