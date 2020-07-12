// Complexity Analysis
// Time: O(M + N), where M and N are the lengths of the given strings.
// Space: O(max(M, N)), to store the final result.


func addBinary(a string, b string) string {
    var aInt big.Int
    var bInt big.Int
    var sum big.Int
    var carry big.Int
    var zero big.Int
    
    aInt.SetString(a, 2)
    bInt.SetString(b, 2)
    
    for bInt.Cmp(&zero) != 0 {
        sum.Xor(&aInt, &bInt)
        carry.And(&aInt, &bInt).Lsh(&carry, 1)
        
        aInt.Set(&sum)
        bInt.Set(&carry)
    }
    
    return aInt.Text(2)
}
