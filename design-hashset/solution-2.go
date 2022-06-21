type MyHashSet struct {
    set [1000001]bool
}


func Constructor() MyHashSet {
    return MyHashSet{}
}


func (this *MyHashSet) Add(key int)  {
    this.set[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    this.set[key] = false
}


func (this *MyHashSet) Contains(key int) bool {
    return this.set[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
