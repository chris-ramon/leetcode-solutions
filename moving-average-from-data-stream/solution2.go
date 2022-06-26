type MovingAverage struct {
    window []int
    size int
}


func Constructor(size int) MovingAverage {
    return MovingAverage{size: size}
}


func (this *MovingAverage) Next(val int) float64 {
    this.window = append(this.window, val)
    if len(this.window) > this.size {
        this.window = this.window[len(this.window)-this.size:]
    }
    sum := 0.0
    for _, num := range this.window {
        sum += float64(num)
    }
    return sum / float64(len(this.window))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */