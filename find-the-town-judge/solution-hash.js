/**
 * @param {number} N
 * @param {number[][]} trust
 * @return {number}
 */
var findJudge = function(N, trust) {
    if(!trust.length && N === 1) return 1;
    if(!trust.length && N > 1) return -1;
    var m = new Map();
    for(var i = 0; i < trust.length; i++) {
        var [person, trustTo] = trust[i];
        if(m.has(person)) {
            const s = m.get(person);
            s.add(trustTo);
            m.set(person, s);
        } else {
            const s = new Set();
            s.add(trustTo);
            m.set(person, s);
        }
        if(!m.has(trustTo)) {
            m.set(trustTo, new Set());
        }
    }
    
    const possibleJudges = [];
    
    for(var [k, v] of m) {
        if(v.size === 0) {
            possibleJudges.push(k);
        }
    }
    
    for(var i = 0; i < possibleJudges.length; i++) {
        const pj = possibleJudges[i];
        var foundJudge = true;
        for (var [k, v] of m) {
            if(k === pj) continue;
            if(!v.has(pj)) {
                foundJudge = false;
            }
        }
        if(foundJudge) {
            return pj;
        }
    }
    
    return -1;
};