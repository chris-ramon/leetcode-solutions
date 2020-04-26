/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
    var p1 = 0; var p2 = 0;
    var visited1 = new Set();
    var visited2 = new Set();
    var result = new Set();
    
    while(nums1.length > p1 || nums2.length > p2) {
        var n1 = p1 < nums1.length ? nums1[p1] : null;
        var n2 = p2 < nums2.length ? nums2[p2] : null;
        
        if(n1 == n2) {
            result.add(n1);
        }
    
        if(visited2.has(n1)) result.add(n1);
        if(visited1.has(n2)) result.add(n2);
    
        if(n1 != null) visited1.add(n1);
        if(n2 != null) visited2.add(n2);
        
        p1 += 1;
        p2 += 1;
    }
    
    return Array.from(result);
};