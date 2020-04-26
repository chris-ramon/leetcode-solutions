func ladderLength(beginWord string, endWord string, wordList []string) int {
    var (
        level int = 1
        queue []string
        words map[string]bool = make(map[string]bool, 0)
    )
    
    for _, w := range wordList {
        words[w] = true
    }
    
    if _, found := words[endWord]; !found {
        return 0
    }
    
    queue = append(queue, beginWord)
    
    
    for len(queue) > 0 {
        size := len(queue)
        
        for i := 0; i < size; i++ {
            newWord := queue[0]
            queue = queue[1:]

            var cwSize = len(newWord)
            
            for k := 0; k < cwSize; k++ {
                for l := 0; l < 26; l++ {
                    
                    cwSplitted := strings.Split(newWord, "")
                    cwSplitted[k] = fmt.Sprintf("%c", 97 + l)  
                    nextWord := strings.Join(cwSplitted, "")

                    if _, found := words[nextWord]; found {
                        if nextWord == endWord {
                            level += 1
                            return level
                        } else {
                            delete(words, nextWord)
                            queue = append(queue, nextWord)
                        }
                    }
                }
            }
        }
        
        level += 1
    }
    
    return 0
}