/**
 * @param {string[]} words
 * @param {string} order
 * @return {boolean}
 */
var isAlienSorted = function(words, order) {
    const dict = new Map();
    order.split("").forEach((o, i) => dict.set(o, i + 1));
    var prev = "";
    for(var i = 0; i < words.length; i++) {
        const word = words[i];
        if(prev === "") {
            prev = word;
            continue;
        }
        for(var j = 0; j < prev.length; j++) {
            const l1 = prev[j];
            const l2 = word[j];
            if(l2 === undefined) return false;
            if(l1 === l2) continue;
            if(dict.get(l1) > dict.get(l2)) return false;
            if(dict.get(l1) < dict.get(l2)) break;
        }
        prev = word;
    }
    return true;
};

// Testcases:
/*
["zello","leetcode"]
"hlabcdefgijkmnopqrstuvwxyz"
["word","world","row"]
"worldabcefghijkmnpqstuvxyz"
["apple","app"]
"abcdefghijklmnopqrstuvwxyz"
*/