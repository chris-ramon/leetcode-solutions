/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    var stack = []; var closeStack = [];
    for(var i = 0; i < s.length; i++) {
        stack.push(s[i]);
    }
    while(stack.length) {
        var c = stack.pop();
        if(c == ")" || c == "}" || c == "]") {
            closeStack.push(c);
            continue;
        }
        if(c == "(" || c == "{" || c == "[") {
            var closingChar = closeStack.pop();
            if(c == "(" && closingChar == ")") {
                continue;
            }
            if(c == "{" && closingChar == "}") {
                continue;
            }
            if(c == "[" && closingChar == "]") {
                continue;
            }
            return false;
        }        
    }
    return closeStack.length === 0;
};