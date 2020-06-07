/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    var result = null;
    var sum = 0;
    for(var i = 0; i < nums.length; i++) {
        var n = nums[i];
        sum = n > sum + n ? n : sum + n;
        result = result === null ? sum : Math.max(result, sum);
    }
    return result;
};