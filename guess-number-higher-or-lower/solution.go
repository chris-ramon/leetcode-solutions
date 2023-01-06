// # Complexity
// - Time complexity: `O(logN)`, where N is the given num.
// - Space complexity: `O(1)`, constant number of variables used.

func guessNumber(n int) int {
	lowerBound := 1
	upperBound := n
	for lowerBound <= upperBound {
		currentGuess := (upperBound + lowerBound) / 2
		switch guess(currentGuess) {
		case 0:
			return currentGuess
		case -1:
			upperBound = currentGuess - 1
		case 1:
			lowerBound = currentGuess + 1
		}
	}
	return -1
}
