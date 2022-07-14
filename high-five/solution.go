// **Complexity Analysis**
// * Time: O(N log N), where N is the size of the given items.
// * Space: O(N), where N is the size of the map used to store the student scores as priority queue.

import "container/heap"

func highFive(items [][]int) [][]int {
	studentsScores := make(map[int]*PriorityQueue)
	for i := 0; i < len(items); i++ {
		id, score := items[i][0], items[i][1]
		if _, found := studentsScores[id]; !found {
			q := NewPriorityQueue(5)
			heap.Init(q)
			studentsScores[id] = q
		}
		studentsScores[id].Append(score)
	}
	studentIds := []int{}
	for id, _ := range studentsScores {
		studentIds = append(studentIds, id)
	}
	sort.Slice(studentIds, func(i, j int) bool {
		return studentIds[i] < studentIds[j]
	})
	result := [][]int{}
	for _, id := range studentIds {
		studentScores := studentsScores[id]
		sum := 0
		size := studentScores.Len()
		for i := 0; i < size; i++ {
			score := heap.Pop(studentScores).(int)
			sum += score
		}
		topFiveAverage := sum / 5
		result = append(result, []int{id, topFiveAverage})
	}
	return result
}

type PriorityQueue struct {
	scores  []int
	maxSize int
}

func (p *PriorityQueue) Push(v interface{}) {
	p.scores = append(p.scores, v.(int))
}

func (p *PriorityQueue) Append(v interface{}) {
	heap.Push(p, v)
	if p.Len() > p.maxSize {
		heap.Pop(p)
	}
}

func (p *PriorityQueue) Pop() interface{} {
	v := p.scores[len(p.scores)-1]
	p.scores = p.scores[:len(p.scores)-1]
	return v
}

func (p *PriorityQueue) Less(i, j int) bool {
	return p.scores[i] < p.scores[j]
}

func (p *PriorityQueue) Swap(i, j int) {
	p.scores[i], p.scores[j] = p.scores[j], p.scores[i]
}

func (p *PriorityQueue) Len() int {
	return len(p.scores)
}

func NewPriorityQueue(maxSize int) *PriorityQueue {
	return &PriorityQueue{maxSize: maxSize}
}
