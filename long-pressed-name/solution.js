/**
 * @param {string} name
 * @param {string} typed
 * @return {boolean}
 */
var isLongPressedName = function(name, typed) {
    var p1 = 0; var p2 = 0;
    while(p1 < name.length || p2 < typed.length) {
        var s1 = p1 < name.length ? name[p1]: null;
        var s2 = p2 < typed.length ? typed[p2]: null;
        
        if(s1 === s2) {
            p1 += 1;
            p2 += 1;
            continue;
        }
        
        if(p1 > 0 && s2 === name[p1 - 1]) {
            p2 += 1;
            continue;
        }
        
        if(s1 !== s2) {
            return false;
        }
    }
    
    return true;
};