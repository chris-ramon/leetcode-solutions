import "strings"

func reverseVowels(s string) string {
    var (
        p1 int = 0
        p2 int = len(s) - 1
        result []byte
    )
    result = make([]byte, len(s))
    for p1 <= p2 {
        var s1 = s[p1];
        var s2 = s[p2];
        if isVowel(s1) && isVowel(s2) {
            result[p1] = s2
            result[p2] = s1
            p1 += 1
            p2 -= 1
            continue
        }
        if !isVowel(s1) {
            result[p1] = s1
            p1 += 1
        }
        if !isVowel(s2) {
            result[p2] = s2
            p2 -= 1
        }
    }
    
    return string(result)
}

func isVowel(vowel byte) bool {
    v := strings.ToLower(string(vowel))
    if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" {
        return true
    }
    return false
}
