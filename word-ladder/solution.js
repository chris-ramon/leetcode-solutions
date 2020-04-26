/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function(beginWord, endWord, wordList) {
    var words = new Set();
    wordList.forEach(w => words.add(w) );
    
    if(!words.has(endWord)) return 0;
    
    var queue = [beginWord];
    var level = 1;
    
    while(queue.length) {
        var size = queue.length;
        
        for (var i = 0; i < size; i++) {
            var word = queue.shift();
            var wSize = word.length;
            
            for(var j = 0; j < wSize; j++) {
                const wordSplitted = word.split("");
                
                for(var k = 0; k < 26; k++) {
                    wordSplitted[j] = String.fromCharCode(97 + k);
                    const nextWord = wordSplitted.join("");
                    if(words.has(nextWord) && nextWord === endWord) {
                        level += 1;
                        return level;
                    }
                    if(words.has(nextWord) && nextWord !== endWord) {
                        queue.push(nextWord);
                        words.delete(nextWord);
                    }
                }
            }
        }
        
        level += 1;
    }
    
    return 0;
};