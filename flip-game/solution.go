/*
Complexity Analysis
- Time: O(n), where n is the size of the current state string.
- Space: O(n), where n is the size of the slice used to store the string.
*/


func generatePossibleNextMoves(currentState string) []string {
	result := []string{}
	for i := range currentState {
		if len(currentState[i:]) >= 2 && currentState[i:i+2] == "++" {
			result = append(result, currentState[:i]+"--"+currentState[i+2:])
		}
	}
	return result
}
