/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    var p1 = m - 1; var p2 = n - 1; var p3 = m + n - 1;
    while(p2 >= 0 || p1 >= 0) {
        var n1 = nums1[p1];
        var n2 = nums2[p2];
        if(p2 < 0) {
            nums1[p3] = n1;
            p1 -= 1;
            p3 -= 1;
        }
        if(p1 < 0) {
            nums1[p3] = n2;
            p2 -= 1;
            p3 -= 1;
        }
        if(n2 > n1) {
            nums1[p3] = n2;
            p2 -= 1;
            p3 -= 1;
            continue;
        }
        if(n2 <= n1) {
            nums1[p3] = n1;
            p1 -= 1;
            p3 -= 1;
            continue;
        }
    }
};
