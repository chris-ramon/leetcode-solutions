// **Complexity Analysis**
// * Time: O(N), where N is the size of the stream given via the constructor.
// * Space: O(N), where N is the size of the stream used to store the inserted values.

type OrderedStream struct {
	ptr    int
	stream []*string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:    1,
		stream: make([]*string, n+1),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = &value
	chunk := []string{}
	if idKey != this.ptr {
		return chunk
	}
	for i := this.ptr; i < len(this.stream); i++ {
		if this.stream[i] == nil {
			this.ptr = i
			break
		}
		chunk = append(chunk, *this.stream[i])
	}
	return chunk
}
