/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function(s) {
    var p1 = 0; var p2 = s.length - 1; var r = Array.from(s);
    var isVowel = (v) => {
        v = v.toLowerCase();
        return v === "a" || v === "e" || v === "i" || v === "o" || v === "u";
    }
    while(p1 < p2) {
        var s1 = s[p1];
        var s2 = s[p2];
        if(isVowel(s1) && isVowel(s2)) {
            r[p1] = s2;
            r[p2] = s1;
            p1 += 1;
            p2 -= 1;
            continue;
        }
        if(!isVowel(s1)) {
            p1 += 1;
        }
        if(!isVowel(s2)) {
            p2 -= 1;
        }

    }
    return r.join("");
};