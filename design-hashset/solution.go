// **Complexity Analysis**
// * Time: O(N/K), where N is the number of possible values and K is the number of buckets.
// * Space: O(N+K), where N is the number of inserted items and K is the number of allocated buckets.

type Node struct {
    Value int
    Next *Node
}

func NewNode(value int, next *Node) *Node {
    return &Node{Value: value, Next: next}
}

type Bucket struct {
    sentinel Node
}

func (b *Bucket) Add(key int) {
    node := NewNode(key, b.sentinel.Next)
    b.sentinel.Next = node
}

func (b *Bucket) Remove(key int) {
    prev := &b.sentinel
    head := b.sentinel.Next
    for head != nil {
        if head.Value == key {
            prev.Next = head.Next
        }
        prev = head
        head = head.Next
    }
}

func (b *Bucket) Contains(key int) bool {
    head := b.sentinel.Next
    for head != nil {
        if head.Value == key {
            return true
        }
        head = head.Next
    }
    return false
}

type MyHashSet struct {
    set [769]Bucket
}

func Constructor() MyHashSet {
    return MyHashSet{}
}

func (this *MyHashSet) hash(key int) int {
    return key % 769
}

func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        this.set[this.hash(key)].Add(key)
    }
}

func (this *MyHashSet) Remove(key int)  {
    this.set[this.hash(key)].Remove(key)
}

func (this *MyHashSet) Contains(key int) bool {
    return this.set[this.hash(key)].Contains(key)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

/*
TestCases
["MyHashSet","add","remove"]
[[],[1],[1]]

["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]

["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains","remove"]
[[], [1], [2], [1], [3], [2], [2], [2], [2],[100]]

["MyHashSet","add","add","contains","contains","add","contains","remove","contains","add","contains"]
[[],[1],[2],[1],[3],[2],[2],[2],[2],[1000000],[1000000]]

["MyHashSet","add","remove","add","remove","remove","add","add","add","add","remove"]
[[],[9],[19],[14],[19],[9],[0],[3],[4],[0],[9]]

["MyHashSet","add","contains","add","contains","remove","add","contains","add","add","add","add","add","add","contains","add","add","add","contains","remove","contains","contains","add","remove","add","remove","add","remove","add","contains","add","add","contains","add","add","add","add","remove","contains","add","contains","add","add","add","remove","remove","add","contains","add","add","contains","remove","add","contains","add","remove","remove","add","contains","add","contains","contains","add","add","remove","remove","add","remove","add","add","add","add","add","add","remove","remove","add","remove","add","add","add","add","contains","add","remove","remove","remove","remove","add","add","add","add","contains","add","add","add","add","add","add","add","add"]
[[],[58],[0],[14],[58],[91],[6],[58],[66],[51],[16],[40],[52],[48],[40],[42],[85],[36],[16],[0],[43],[6],[3],[25],[99],[66],[60],[58],[97],[3],[35],[65],[40],[41],[10],[37],[65],[37],[40],[28],[60],[30],[63],[76],[90],[3],[43],[81],[61],[39],[75],[10],[55],[92],[71],[2],[20],[7],[55],[88],[39],[97],[44],[1],[51],[89],[37],[19],[3],[13],[11],[68],[18],[17],[41],[87],[48],[43],[68],[80],[35],[2],[17],[71],[90],[83],[42],[88],[16],[37],[33],[66],[59],[6],[79],[77],[14],[69],[36],[21],[40]]
*/
