/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @param {number[]} arr3
 * @return {number[]}
 */
var arraysIntersection = function(arr1, arr2, arr3) {
    var p1 = 0; var p2 = 0; var p3 = 0; var result = [];
    while(p1 < arr1.length && p2 < arr2.length && p3 < arr3.length) {
        var n1 = arr1[p1];
        var n2 = arr2[p2];
        var n3 = arr3[p3];
        
        if(n1 === n2 && n2 === n3) {
            result.push(n1);
            p1 += 1;
            p2 += 1;
            p3 += 1;
            continue;
        }
        
        var max = Math.max(n1, n2, n3);
        
        if(n1 < max) {
            p1 += 1;
        }
        if(n2 < max) {
            p2 += 1;
        }
        if(n3 < max) {
            p3 += 1;
        }
    }
    return result;
};