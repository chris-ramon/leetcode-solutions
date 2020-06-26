/**
 * @param {number} N
 * @param {number[][]} trust
 * @return {number}
 */
var findJudge = function(N, trust) {
    if(N === 1 && !trust.length) return 1;
    if(N > 1 && !trust.length) return -1;
    
    const inDegree = new Map();
    const outDegree = new Map();
    
    for(var i = 0; i < trust.length; i++) {
        const [truster, trusted] = trust[i];
        if(inDegree.has(trusted)) {
            inDegree.set(trusted, inDegree.get(trusted) + 1);
        } else {
            inDegree.set(trusted, 1);
        }
        if(outDegree.has(truster)) {
            outDegree.set(truster, outDegree.get(truster) + 1);
        } else {
            outDegree.set(truster, 1);
        }
    }
    
    for(var i = 1; i <= N; i++) {
        if(inDegree.get(i) === N - 1 && (!outDegree.has(i) || outDegree.get(i) === 0)) {
            return i;
        }
    }
    
    return -1;

};