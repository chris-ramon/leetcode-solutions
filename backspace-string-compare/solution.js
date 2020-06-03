/**
 * @param {string} S
 * @param {string} T
 * @return {boolean}
 */
var backspaceCompare = function(S, T) {
    var p1 = S.length - 1; var p2 = T.length - 1;
    while(p1 >= 0 || p2 >= 0) {
        var skipS = 0;
        while(p1 >= 0) {
            var n1 = S[p1];
            if(n1 !== "#" && skipS === 0) {
                break;
            }
            if(n1 === "#") {
                skipS += 1;
                p1 -= 1;
                continue;
            }
            if(skipS > 0) {
                skipS -= 1;
                p1 -= 1;
            }
        }
        
        var skipT = 0;
        while(p2 >= 0) {
            var n2 = T[p2];
            if(n2 !== "#" && skipT === 0) {
                break;
            }
            if(n2 === "#") {
                skipT += 1;
                p2 -= 1;
                continue;
            }
            if(skipT > 0) {
                skipT -= 1;
                p2 -= 1;
            }
        }
        
        if(S[p1] !== T[p2]) {
            return false;
        }
        
        p1 -= 1;
        p2 -= 1;
    }
    
    return true;
};