type Flight struct {
	From  int
	To    int
	Price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	flightsMap := make(map[int][]Flight)
	for i := 0; i < len(flights); i++ {
		from, to, price := flights[i][0], flights[i][1], flights[i][2]
		flightsMap[from] = append(flightsMap[from], Flight{From: from, To: to, Price: price})
	}
	result := -1
	dfs(src, flightsMap, -1, k, &result, dst, 0)
	return result
}

func dfs(src int, flightsMap map[int][]Flight, stops int, maxStops int, result *int, dst int, subPrice int) {
	if stops >= maxStops {
		return
	}
	for _, flight := range flightsMap[src] {
		totalPrice := subPrice + flight.Price
		if flight.To == dst && (*result == -1 || totalPrice < *result) {
			*result = totalPrice
		}
		dfs(flight.To, flightsMap, stops+1, maxStops, result, dst, totalPrice)
	}
}

/*
TestCases

4
[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]
0
3
1

3
[[0,1,100],[1,2,100],[0,2,500]]
0
2
1

3
[[0,1,100],[1,2,100],[0,2,500]]
0
2
0

3
[[0,1,2],[1,2,1],[2,0,10]]
1
2
1

4
[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]]
0
3
1

10
[[3,4,4],[2,5,6],[4,7,10],[9,6,5],[7,4,4],[6,2,10],[6,8,6],[7,9,4],[1,5,4],[1,0,4],[9,7,3],[7,0,5],[6,5,8],[1,7,6],[4,0,9],[5,9,1],[8,7,3],[1,2,6],[4,1,5],[5,2,4],[1,9,1],[7,8,10],[0,4,2],[7,2,8]]
6
0
7

*/
