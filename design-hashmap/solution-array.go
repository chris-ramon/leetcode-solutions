type MyHashMap struct {
    h [1000001]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    m := MyHashMap{}
    for i := 0; i < 1000001; i++ {
        m.h[i] = -1
    }
    return m
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    this.h[key] = value
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    return this.h[key]
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    this.h[key] = -1
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