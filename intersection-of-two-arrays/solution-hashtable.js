/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
    var result = new Set();
    var nums = new Set();
    nums1.forEach(n => nums.add(n));
    nums2.forEach(n => {
       if(nums.has(n) && !result.has(n)) result.add(n);
    });
    return Array.from(result);
};