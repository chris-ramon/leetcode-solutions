/**
 * @param {string[]} logs
 * @return {string[]}
 */
var reorderLogFiles = function(logs) {
    var letterLogs = [];
    var digitLogs = [];
    
    logs.forEach((log) => {
        const logSplitted = log.split(" ");
        const identifier = logSplitted.shift();
        const body = logSplitted.join(" ");
        if(isNaN(logSplitted.join(""))) {
            letterLogs.push({identifier, body});
        } else {
            digitLogs.push({identifier, body});
        }
    });
    
    letterLogs.sort((a, b) => {
        if(a.body == b.body) {
            return a.identifier.localeCompare(b.identifier);
        }
        return a.body.localeCompare(b.body); 
    });
    
    var result = [];
    letterLogs.forEach(ll => result.push(ll.identifier + " " + ll.body));
    digitLogs.forEach(dl => result.push(dl.identifier + " " + dl.body));
    
    return result;
};