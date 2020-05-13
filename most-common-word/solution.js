/**
 * @param {string} paragraph
 * @param {string[]} banned
 * @return {string}
 */
var mostCommonWord = function(paragraph, banned) {
    var mostCommonWord = "";
    var mostCommonWordTotal = 0;
    var words = paragraph.toLowerCase().split(" ");
    var wordRegExp = /[\W+]?(\w+)[\W*]?/;
    const bannedWordsMap = new Set();
    banned.forEach(b => bannedWordsMap.add(b));
    const wordsMap = new Map();
    words.forEach(w => {
        const wMatch = w.match(wordRegExp);
        if(!w || !w.length) return;
        const word = wMatch[1];
        if(wordsMap.has(word)) {
            wordsMap.set(word, wordsMap.get(word) + 1);
        } else {
            wordsMap.set(word, 1);
        }
        if(!bannedWordsMap.has(word) && wordsMap.get(word) > mostCommonWordTotal) {
            mostCommonWordTotal = wordsMap.get(word);
            mostCommonWord = word;
        }
    });
    return mostCommonWord;
};