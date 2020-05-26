/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    if(needle === "") return 0;
    var p1 = 0;
    var hl = haystack.length;
    var nl = needle.length;
    
    while(p1 <= (hl - nl)) {
        
        var p2 = 0;
        
        while(p2 <= nl - 1 && (haystack[p1] === needle[p2])) {
            p1 += 1;
            p2 += 1;
        }
        
        if(p2 === nl) {
            return p1 - p2;
        }
        
        p1 = p1 - p2 + 1;
    }
    
    return -1;
};
