/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 * hello
 *    ^^
 */
var strStr = function(haystack, needle) {
    var hl = haystack.length;
    var nl = needle.length;
    
    // O((m - n) * n), where m is haystack size and n is needle size.
    // We compute the substring of needle's size: O(n).
    for(var i = 0; i < hl - nl + 1; i++) {
        if(haystack.substring(i, i + nl) === needle) {
            return i;
        }
    }
    
    return -1;
};
