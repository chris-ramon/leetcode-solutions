// **Complexity Analysis**
// * Time: O(1), constant time as number of messages grows.
// * Space: O(N), where N is the size of the map used to store the messages.

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{m: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	t, found := this.m[message]
	if !found || timestamp >= t {
		this.m[message] = 10 + timestamp
		return true
	}
	return false
}
