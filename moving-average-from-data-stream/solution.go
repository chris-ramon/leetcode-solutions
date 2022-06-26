type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

type MovingAverage struct {
	maxSize     int
	currentSize int
	currentSum  int
	head        *Node
	tail        *Node
}

func Constructor(size int) MovingAverage {
	return MovingAverage{maxSize: size}
}

func (this *MovingAverage) Next(val int) float64 {
	node := NewNode(val)
	if this.head == nil {
		this.head = node
	} else {
		temp := this.head
		temp.Prev = node
		node.Next = temp
		this.head = node
	}
	if this.tail == nil {
		this.tail = node
	}

	this.currentSize += 1
	tailVal := 0

	if this.currentSize > this.maxSize {
		temp := this.tail
		this.tail = temp.Prev
		this.currentSize -= 1
		tailVal = temp.Val
	}

	this.currentSum = (this.currentSum - tailVal) + val
	return float64(this.currentSum) / float64(this.currentSize)
}