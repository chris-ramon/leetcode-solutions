/*
Time Complexity:
- Time: O(n): where n is the size of the queue. 
- Space: O(n): where n is the size of the queue.
*/

// @lc code=start
type MyQueue struct {
	s1 []int
	s2 []int
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.s1 = append(this.s1, x)
}


func (this *MyQueue) Pop() int {
	last := this.s1[0]
	this.s1 = this.s1[1:]
	return last
}


func (this *MyQueue) Peek() int {
	last := this.s1[0]
	return last
}


func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0
}

/*
TestCases:
["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]\n
["MyQueue","push","push","push","pop","empty"]\n[[],[1],[2],[3],[],[]]\n
["MyQueue","push","push","pop","pop","empty"]\n[[],[1],[2],[],[],[]]\n
*/