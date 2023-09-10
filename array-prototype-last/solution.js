// Complexity Analysis
// Time: O(1), constant time operation.
// Space: O(1), constant number of variables used.

Array.prototype.last = function() {
    if(this.length === 0) {
        return -1;
    }
    return this[this.length-1];
};

/**
 * const arr = [1, 2, 3];
 * arr.last(); // 3
 */
