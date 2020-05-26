/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    var p1 = 0;
    for(var i = 0; i < nums.length; i++) {
        var n = nums[i];
        if(n !== 0) {
            nums[p1] = n;
            p1 += 1;
        }
    }
    for(var i = p1; i < nums.length; i++) {
        nums[i] = 0;
    }
};