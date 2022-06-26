func countBinarySubstrings(s string) int {
    result := 0.0
    prev := 0.0
    current := 1.0
    for i := 1; i < len(s); i++ {
        if s[i-1] != s[i] {
            result += math.Min(prev, current)
            prev = current
            current = 1.0
        } else {
            current += 1.0
        }
    }
    return int(result + math.Min(prev, current))
}