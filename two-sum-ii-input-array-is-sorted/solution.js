/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    if(!numbers || !numbers.length) return [];
    var p1 = 0; var p2 = 1;
    while(p1 < numbers.length - 1) {
        var n1 = numbers[p1];
        var n2 = numbers[p2];
        if(n1 + n2 === target) {
            return [p1 + 1, p2 + 1];
        }
        p2 += 1;
        if(p2 === numbers.length) {
            p1 += 1;
            p2 = p1 + 1;
        }
    }
    return [];
};