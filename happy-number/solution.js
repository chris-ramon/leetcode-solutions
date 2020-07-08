/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    var slow = n;
    var fast = next(n);
    
    while(fast != 1 && slow != fast) {
        slow = next(slow);
        fast = next(next(fast));
    }
    
    return fast === 1;
};

var next = function(n) {
    var sum = 0;
    while(n > 0) {
        var d = n % 10
        sum += d * d;
        n = parseInt(n / 10);
    }
    return sum;
};

// Testcases:
/*
19
100
33
0
1
*/