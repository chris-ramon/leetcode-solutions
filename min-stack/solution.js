/**
 * initialize your data structure here.
 */
var MinStack = function() {
    this.stack = []; 
    this.minStack = [];
    this.addToMinStack = (x) => {
        var min = this.minStack[this.minStack.length - 1];
        if(min !== undefined && x > min) return;
        this.minStack.push(x);
    };
    this.removeFromMinStack = (x) => {
        var min = this.minStack[this.minStack.length - 1];
        if(x === min) this.minStack.pop();
    };
};

/** 
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
    this.stack.push(x);
    this.addToMinStack(x);
    return;
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
    var x = this.stack.pop();
    this.removeFromMinStack(x);
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
    return this.stack[this.stack.length - 1];
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
    return this.minStack[this.minStack.length - 1];
};

/** 
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */