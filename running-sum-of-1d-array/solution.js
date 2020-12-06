// Complexity Analysis
// Time: `O(n)` where n is the number of given nums.
// Space: `O(n)`, where n is the size of the result array.

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function (nums) {
  var sum = 0;
  var result = [];
  for (var i = 0; i < nums.length; i++) {
    result[i] = nums[i] + sum;
    sum += nums[i];
  }
  return result;
};
