// **Complextiy Analysis**
// * Time: O(1), constant runtime for the Queue methods: Push, Pop, Top & Empty.
// * Space: O(N), where N is the number of queues used to store the Stack values.


import "container/list"

type MyStack struct {
	q *list.List
}

func Constructor() MyStack {
	return MyStack{q: list.New()}
}

func (this *MyStack) Push(x int) {
	q := list.New()
	q.PushBack(x)
	q.PushBack(this.q)
	this.q = q
}

func (this *MyStack) Pop() int {
	el := this.q.Front()
	this.q.Remove(el)
	this.q = this.q.Front().Value.(*list.List)
	return el.Value.(int)
}

func (this *MyStack) Top() int {
	return this.q.Front().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}

/*
TestCases
["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]\n["MyStack","push","push","pop","pop","empty"]\n[[],[1],[2],[],[],[]]\n["MyStack","push","push","pop","pop","push","empty"]\n[[],[1],[2],[],[],[10],[]]
*/
