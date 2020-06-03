/**
 * @param {number[]} A
 * @return {number[]}
 */
var sortedSquares = function(A) {
    var p1 = null; var p2 = null; var p3 = 0; var result = [];
    for(var i = 0; i < A.length; i++) {
        if(A[i] < 0) {
            p1 = i;
        }
        if(A[i] >= 0) {
            p2 = i;
            break;
        }
    }
    while((p1 != null && p1 >= 0) || (p2 != null && p2 < A.length)) {
        var n1 = p1 >= 0 ? A[p1]: null;
        var n2 = p2 < A.length ? A[p2]: null;
        if(n1 != null && n2 != null && Math.abs(n1) < Math.abs(n2)) {
            result.push(Math.pow(n1, 2));
            p1 -= 1;
            continue;
        }
        if(n1 != null && n2 != null && Math.abs(n1) > Math.abs(n2)) {
            result.push(Math.pow(n2, 2));
            p2 += 1;
            continue;
        }
        if(n1 != null) {
            result.push(Math.pow(n1, 2));
            p1 -= 1;
            continue;
        }
        if(n2 != null) {
            result.push(Math.pow(n2, 2));
            p2 += 1;
        }
    }
    return result;
};