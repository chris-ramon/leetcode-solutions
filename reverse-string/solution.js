/**
 * @param {character[]} s
 * @return {void} Do not return anything, modify s in-place instead.
 */
var reverseString = function(s) {
    if(!s || !s.length) return;
    var p1 = 0; var p2 = s.length - 1;
    while(p1 <= p2) {
        var s1 = s[p1];
        var s2 = s[p2];
        s[p1] = s2;
        s[p2] = s1;
        p1 += 1;
        p2 -= 1;
    }
};