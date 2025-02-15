// **Complexity Analysis**
// - Time: O(N^2 * log N), where N is the number of cities within the flights.
// - Space: O(N^2), where N is the number of cities stored in the heap.

import "container/heap"

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
	pq := NewPriorityQueue()
	heap.Init(pq)
	heap.Push(pq, &Node{ID: src, Cost: 0, Depth: 0})
	bestVisited := make(map[int]int)
	for pq.Len() > 0 {
		n := heap.Pop(pq).(*Node)
		if depth, found := bestVisited[n.ID]; found && depth <= n.Depth {
			continue
		}
		if n.ID == dst {
			return n.Cost
		}
		bestVisited[n.ID] = n.Depth
		if n.Depth <= k {
			for _, flight := range flightsMap[n.ID] {
				cost := flight.Price + n.Cost
				depth, found := bestVisited[flight.To]
				if !found || n.Depth+1 < depth {
					heap.Push(pq, &Node{ID: flight.To, Cost: cost, Depth: n.Depth + 1})
				}
			}
		}
	}
	return result
}

type Node struct {
	ID    int
	Cost  int
	Depth int
}

type PriorityQueue struct{ queue []*Node }

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq *PriorityQueue) Len() int           { return len(pq.queue) }
func (pq *PriorityQueue) Less(i, j int) bool { return pq.queue[i].Cost < pq.queue[j].Cost }
func (pq *PriorityQueue) Swap(i, j int)      { pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	pq.queue = append(pq.queue, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[:len(pq.queue)-1]
	return n
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

17
[[0,12,28],[5,6,39],[8,6,59],[13,15,7],[13,12,38],[10,12,35],[15,3,23],[7,11,26],[9,4,65],[10,2,38],[4,7,7],[14,15,31],[2,12,44],[8,10,34],[13,6,29],[5,14,89],[11,16,13],[7,3,46],[10,15,19],[12,4,58],[13,16,11],[16,4,76],[2,0,12],[15,0,22],[16,12,13],[7,1,29],[7,14,100],[16,1,14],[9,6,74],[11,1,73],[2,11,60],[10,11,85],[2,5,49],[3,4,17],[4,9,77],[16,3,47],[15,6,78],[14,1,90],[10,5,95],[1,11,30],[11,0,37],[10,4,86],[0,8,57],[6,14,68],[16,8,3],[13,0,65],[2,13,6],[5,13,5],[8,11,31],[6,10,20],[6,2,33],[9,1,3],[14,9,58],[12,3,19],[11,2,74],[12,14,48],[16,11,100],[3,12,38],[12,13,77],[10,9,99],[15,13,98],[15,12,71],[1,4,28],[7,0,83],[3,5,100],[8,9,14],[15,11,57],[3,6,65],[1,3,45],[14,7,74],[2,10,39],[4,8,73],[13,5,77],[10,0,43],[12,9,92],[8,2,26],[1,7,7],[9,12,10],[13,11,64],[8,13,80],[6,12,74],[9,7,35],[0,15,48],[3,7,87],[16,9,42],[5,16,64],[4,5,65],[15,14,70],[12,0,13],[16,14,52],[3,10,80],[14,11,85],[15,2,77],[4,11,19],[2,7,49],[10,7,78],[14,6,84],[13,7,50],[11,6,75],[5,10,46],[13,8,43],[9,10,49],[7,12,64],[0,10,76],[5,9,77],[8,3,28],[11,9,28],[12,16,87],[12,6,24],[9,15,94],[5,7,77],[4,10,18],[7,2,11],[9,5,41]]
13
4
13

13
[[11,12,74],[1,8,91],[4,6,13],[7,6,39],[5,12,8],[0,12,54],[8,4,32],[0,11,4],[4,0,91],[11,7,64],[6,3,88],[8,5,80],[11,10,91],[10,0,60],[8,7,92],[12,6,78],[6,2,8],[4,3,54],[3,11,76],[3,12,23],[11,6,79],[6,12,36],[2,11,100],[2,5,49],[7,0,17],[5,8,95],[3,9,98],[8,10,61],[2,12,38],[5,7,58],[9,4,37],[8,6,79],[9,0,1],[2,3,12],[7,10,7],[12,10,52],[7,2,68],[12,2,100],[6,9,53],[7,4,90],[0,5,43],[11,2,52],[11,8,50],[12,4,38],[7,9,94],[2,7,38],[3,7,88],[9,12,20],[12,0,26],[10,5,38],[12,8,50],[0,2,77],[11,0,13],[9,10,76],[2,6,67],[5,6,34],[9,7,62],[5,3,67]]
10
1
10
*/
