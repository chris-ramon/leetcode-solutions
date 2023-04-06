/*
Complexity Analysis
- Time: O(n), where n is the size of file read.
- Time: O(n), where n is the size of data buffer.
*/
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		var data []byte
		localBuff := make([]byte, 4)
		totalLength := 0
		readLength := 4
		for readLength == 4 {
			readLength = read4(localBuff)
			data = append(data, localBuff[:readLength]...)
			totalLength += readLength
		}
		return copy(buf, data[:min(totalLength, n)])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
