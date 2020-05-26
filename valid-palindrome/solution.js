/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    if(s === "") return true;
    var p1 = 0; var p2 = s.length - 1;
    
    while(p2 >= p1) {
        var s1 = s[p1].toLowerCase();
        var s2 = s[p2].toLowerCase();
        if(/\W+/.test(s1)) {
            p1 += 1;
            continue;
        }
        if(/\W+/.test(s2)) {
            p2 -= 1;
            continue;
        }
        if(s1 !== s2) {
            return false;
        }
        p1 += 1;
        p2 -= 1;
    }
    
    return true;
};