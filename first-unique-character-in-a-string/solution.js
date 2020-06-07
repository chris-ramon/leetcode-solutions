/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function(s) {
    if(!s || s === "") return -1;
    var m = new Map();
    for(var i = 0; i < s.length; i++) {
        var c = s[i];
        if(!m.has(c)) {
            m.set(c, 1);
            continue;
        }
        m.set(c, m.get(c) + 1);
    }
    for(var i = 0; i < s.length; i++) {
        var c = s[i];
        if(m.get(c) === 1) {
            return i;
        }
    }
    return -1;
};