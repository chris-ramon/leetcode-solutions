var Node = function(x) {
    this.value = x
    this.next = x
    this.prev = x
}

var MaxStack = function() {
    this.tail = null
};

/** 
 * @param {number} x
 * @return {void}
 */
MaxStack.prototype.push = function(x) {
    const node = new Node(x)
    if(!this.tail) {
        this.tail = node
        return
    }
    
    const temp = this.tail
    temp.next = node
    node.prev = temp
    
    this.tail = node
};

/**
 * @return {number}
 */
MaxStack.prototype.pop = function() {
    const temp = this.tail
    
    temp.next = null
    this.tail = temp.prev
    
    return temp.value
};

/**
 * @return {number}
 */
MaxStack.prototype.top = function() {
    return this.tail.value
};

/**
 * @return {number}
 */
MaxStack.prototype.peekMax = function() {
    let node = this.tail
    let max = null
    while(node) {
        if(max === null || node.value > max) {
            max = node.value
        }
        node = node.prev
    }
    return max
};

/**
 * @return {number}
 */
MaxStack.prototype.popMax = function() {
    let node = this.tail
    let max = new Node(null)
    while(node) {
        if(max.value === null || node.value > max.value) {
            max = node
        }
        node = node.prev
    }
    
    let temp = max
    
    if(this.tail == max) {
        this.tail = max.prev
    }
    
    if(temp.prev) {
        temp.prev.next = temp.next
    }
    if(temp.next) {
        temp.next.prev = temp.prev
    }
    
    return temp.value
};

/** 
 * Your MaxStack object will be instantiated and called as such:
 * var obj = new MaxStack()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.peekMax()
 * var param_5 = obj.popMax()
 */

/**
TestCases:

["MaxStack","push","push","push","top","popMax","top","peekMax","pop","top"]
[[],[5],[1],[5],[],[],[],[],[],[]]

["MaxStack","push","popMax","push","push","push","pop","peekMax","push","pop","pop","push","peekMax","peekMax","push","peekMax","push","peekMax"]
[[],[-2],[],[-45],[-82],[29],[],[],[40],[],[],[66],[],[],[-61],[],[98],[]]

*/