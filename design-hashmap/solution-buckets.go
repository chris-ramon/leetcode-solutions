type MyHashMap struct {
    buckets [][][]int
    mod int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    var mod int = 2069
    
    m := MyHashMap{
        buckets: make([][][]int, mod),
        mod: mod,
    }
    
    return m
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    hash := key % this.mod
    if len(this.buckets[hash]) == 0 {
        this.buckets[hash] = [][]int{[]int{key, value}}
    } else {
        var found bool
        for i, space := range this.buckets[hash] {
            k, _ := space[0], space[1]
            if k == key {
                this.buckets[hash][i] = []int{key, value}
                found = true
                break
            }
        }
        if !found {
            this.buckets[hash] = append(this.buckets[hash], []int{key, value})
        }
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    hash := key % this.mod
    for _, space := range this.buckets[hash] {
        k, v := space[0], space[1]
        if k == key {
            return v
        }
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    hash := key % this.mod
    for i, space := range this.buckets[hash] {
        k, _ := space[0], space[1]
        if k == key {
            this.buckets[hash] = append(this.buckets[hash][:i], this.buckets[hash][i+1:]...)
            break
        }
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

// Testcases:
/*
["MyHashMap","put","put","get","get","put","get", "remove", "get"]
[[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]
["MyHashMap","put","remove", "get"]
[[],[1,1],[1],[1]]
["MyHashMap","put","remove", "get"]
[[],[0,0],[1],[1]]
["MyHashMap","put","remove", "get"]
[[],[0,-1],[1],[1]]
["MyHashMap","put","remove", "get"]
[[],[1000000,1000000],[1000000],[1000000]]
["MyHashMap","remove","put","put","put","put","put","put","get","put","put"]
[[],[2],[3,11],[4,13],[15,6],[6,15],[8,8],[11,0],[11],[1,10],[12,14]]
*/