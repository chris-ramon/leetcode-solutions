// **Complexity Analysis**
// * Time: O(N), where N is the size of the moves.
// * Space: O(N), where N is the size of the maps used to store the moves.

func tictactoe(moves [][]int) string {
	playerA := map[int]int{}
	playerB := map[int]int{}
	for i := 0; i < len(moves); i++ {
		currentPlayer := map[int]int{}
		if i%2 == 0 {
			currentPlayer = playerA
		} else {
			currentPlayer = playerB
		}
		r, c := moves[i][0], moves[i][1]
		currentPlayer[r]++
		currentPlayer[c+3]++
		if r == c {
			currentPlayer[6]++
		}
		if r == 2-c {
			currentPlayer[7]++
		}
	}
	for i := 0; i < 8; i++ {
		if playerA[i] == 3 {
			return "A"
		}
		if playerB[i] == 3 {
			return "B"
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
