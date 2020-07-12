// Complexity Analysis
// Time: O(max(M, N)) where M and N are the given strings length. 
// Space: O(max(M, N)) to keep the result string.


func addBinary(a string, b string) string {
    var p1 int = len(a) - 1
    var p2 int = len(b) - 1
    var result string 
    var carry int 
    
    for p1 >= 0 || p2 >= 0 {
        var i1 int
        var i2 int
        var sum int = carry
        if p1 >= 0 {
            i1 = int(a[p1] - '0') 
        } 
        if p2 >= 0 {
            i2 = int(b[p2] - '0')
        }
        sum += i1 + i2
        switch sum {
            case 0:
                result = "0" + result
                carry = 0
            case 1:
                result = "1" + result
                carry = 0
            case 2:
                result = "0" + result
                carry = 1
            case 3:
                result = "1" + result
                carry = 1
        }
        p1 -= 1
        p2 -= 1
    }
    
    if carry == 1 {
        result = "1" + result
    }
    
    return result
}

// Testcases:
/*
"11"
"1"
"1010"
"1011"
"1"
"0"
"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
"1"
"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000"
*/
