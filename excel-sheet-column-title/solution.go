// # Complexity
// - Time complexity: `O(N)`, where N is the given number.
// - Space complexity: `O(N)`, where N is the given number, used to store the result.

func convertToTitle(columnNumber int) string {
	letters := []byte{}
	currentNumber := columnNumber
	for currentNumber > 0 {
		currentNumber--
		i := currentNumber
		i %= 26
		letters = append(letters, byte('A'+i))
		currentNumber /= 26
	}
	for i, j := 0, len(letters)-1; i <= j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}
	return string(letters)
}
