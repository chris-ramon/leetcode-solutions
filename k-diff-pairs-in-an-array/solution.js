/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findPairs = function(nums, k) {
    var result = 0;
    if(!nums || !nums.length || nums.length === 1) return result;
    var visited = new Set();
    var visitedKey = (n1, n2) => n1 + ":" + n2;
    var p1 = 0; var p2 = 1;
    while(p1 < nums.length - 1) {
        var n1 = nums[p1];
        var n2 = nums[p2];
        if(Math.abs(n1-n2) === k && !visited.has(visitedKey(n1, n2)) && !visited.has(visitedKey(n2, n1))) {
            result += 1;
            visited.add(visitedKey(n1, n2), true);
            visited.add(visitedKey(n2, n1), true);
        }
        p2 += 1;
        if(p2 === nums.length) {
            p1 += 1;
            p2 = p1 + 1;
        }

    }
    return result;
};