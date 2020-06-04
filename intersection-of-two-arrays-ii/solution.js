/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function(nums1, nums2) {
    var result = [];
    var m = new Map();
    nums1.forEach(n => {
        if(!m.has(n)) {
            m.set(n, 1);
            return;
        }
        m.set(n, m.get(n) + 1);
    });
    nums2.forEach(n => {
        if(!m.has(n) || m.get(n) <= 0) {
            return;
        }
        m.set(n, m.get(n) - 1);
        result.push(n);
    });
    return result;
};