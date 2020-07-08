/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function(num1, num2) {
    var p1 = num1.length - 1;
    var p2 = num2.length - 1;
    var carry = false;
    var result = [];
    
    while(p1 >= 0 || p2 >= 0) {
        var n1 = p1 >= 0 ? parseInt(num1[p1]) : 0;
        var n2 = p2 >= 0 ? parseInt(num2[p2]) : 0;
        
        var sum = carry ? n1 + n2 + 1 : n1 + n2;
        carry = sum >= 10;
        
        var remain = sum % 10;
        
        result.push(remain);
        
        p1 -= 1;
        p2 -= 1;
    }
    
    if(carry) {
        result.push(1);
    }
    
    return result.reverse().join("");
};

// Testcases:
/* 
"333"
"22"
"5"
"5"
"6"
"6"
"20"
"1"
*/