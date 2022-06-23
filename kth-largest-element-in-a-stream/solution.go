// Complexity Analysis
// * Time: O(N * Log N), where N is the size of given nums.
// * Space: O(N), where N is the size of given nums, max size of the priority queue.

import "container/heap"

type KthLargest struct {
	k int
	q *PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	q := NewPriorityQueue()
	kthLargest := KthLargest{k: k, q: q}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.q, val)
	for this.q.Len() > this.k {
		heap.Pop(this.q)
	}
	return this.q.Peek()
}

type PriorityQueue struct {
	q []int
}

func (p *PriorityQueue) Push(v interface{}) {
	p.q = append(p.q, v.(int))
}
func (p *PriorityQueue) Pop() interface{} {
	i := p.q[len(p.q)-1]
	p.q = p.q[:len(p.q)-1]
	return i
}
func (p *PriorityQueue) Len() int           { return len(p.q) }
func (p *PriorityQueue) Swap(i, j int)      { p.q[i], p.q[j] = p.q[j], p.q[i] }
func (p *PriorityQueue) Less(i, j int) bool { return p.q[i] < p.q[j] }
func (p *PriorityQueue) Peek() int          { return p.q[0] }
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}
