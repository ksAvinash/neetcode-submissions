type MyHashSet struct {
	items []int
}

func Constructor() MyHashSet {
    return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
    for i, v := range this.items {
		if v == key {
			this.items[i] = key
			return
		}
	}
	this.items = append(this.items, key)
}

func (this *MyHashSet) Remove(key int) {
    for i, v := range this.items {
		if v == key {
			this.items = append(this.items[0:i], this.items[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, v := range this.items {
		if v == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 