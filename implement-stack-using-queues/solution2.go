import "container/list"

type MyStack struct {
	q *list.List
}


func Constructor() MyStack {
	return MyStack{
		q: list.New(),
	}
}


func (this *MyStack) Push(x int)  {
	this.q.PushBack(x)
}


func (this *MyStack) Pop() int {
	top := this.top()
	this.q.Remove(top)
	return top.Value.(int)
}


func (this *MyStack) Top() int {
	return this.top().Value.(int)
}


func (this *MyStack) top() *list.Element {
	for temp := this.q.Front(); temp != nil; temp = temp.Next() {
		if temp.Next() == nil {
			return temp
		}
	}
	return nil
}


func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

/*
TestCases
["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]
["MyStack","push","push","pop","pop","empty"]\n[[],[1],[2],[],[],[]]
["MyStack","push","push","pop","pop","push","empty"]\n[[],[1],[2],[],[],[10],[]]
*/
