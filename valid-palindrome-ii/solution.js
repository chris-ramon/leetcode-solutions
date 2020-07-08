/**
 * @param {string} s
 * @return {boolean}
 */
var validPalindrome = function(s) {
    if(!s || s === "") return true;
    var p1 = 0;
    var p2 = s.length - 1;
    var retry = 0;
    
    while(p1 <= p2) {
        var s1 = s[p1];
        var s2 = s[p2];
        
        if(s1 !== s2 && retry === 2) {
            return false;
        }
        
        if(s1 !== s2 && retry === 0) {
            p2 -= 1;
            retry += 1;
            continue;
        }
        
        if(s1 !== s2 && retry === 1) {
            p2 += 1;
            p1 += 1;
            retry += 1;
            continue;
        }
        
        p1 += 1;
        p2 -= 1;
    }
    
    return true;
};

// Testcases:
/*
"aba"
"abca"
"xa"
"a"
"aczybxzca"
*/