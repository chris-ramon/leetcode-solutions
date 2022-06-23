// **Complexity Analysis**
// - Time: O(N * M), where N x M is the matrix size.
// - Space: O(N * Log N), where N is the number of rows from the matrix used as priority queue max size.


import "container/heap"

type Row struct {
	Count int
	Idx   int
}

func kWeakestRows(mat [][]int, k int) []int {
	q := NewPriorityQueue()
	heap.Init(q)
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				count += 1
			}
		}
		heap.Push(q, &Row{Count: count, Idx: i})
	}
	kWeakest := []*Row{}
	for i := 0; i < k; i++ {
		r := heap.Pop(q).(*Row)
		kWeakest = append(kWeakest, r)
	}
	result := []int{}
	for _, kWeak := range kWeakest {
		result = append(result, kWeak.Idx)
	}
	return result
}

type PriorityQueue struct {
	q []*Row
}

func (p *PriorityQueue) Push(v interface{}) { p.q = append(p.q, v.(*Row)) }
func (p *PriorityQueue) Pop() interface{} {
	i := p.q[len(p.q)-1]
	p.q = p.q[:len(p.q)-1]
	return i
}
func (p *PriorityQueue) Len() int      { return len(p.q) }
func (p *PriorityQueue) Swap(i, j int) { p.q[i], p.q[j] = p.q[j], p.q[i] }
func (p *PriorityQueue) Less(i, j int) bool {
	if p.q[i].Count == p.q[j].Count {
		return p.q[i].Idx < p.q[j].Idx
	}
	return p.q[i].Count < p.q[j].Count
}
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}
