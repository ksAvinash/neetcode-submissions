type MyHashMap struct {
	items [][]int
}

func Constructor() MyHashMap {
    return MyHashMap{}
}

func (this *MyHashMap) Put(key int, value int) {
    for i, v := range this.items {
		if v[0] == key {
			this.items[i] = []int{key, value}
			return
		}
	}
	this.items = append(this.items, []int{key, value})
}

func (this *MyHashMap) Get(key int) int {
    for _, v := range this.items {
		if v[0] == key {
			return v[1]
		}
	}
	return -1
}


func (this *MyHashMap) Remove(key int) {
	pos := -1
    for i, v := range this.items {
		if v[0] == key {
			pos = i
			break
		}
	}
	if pos >= 0 {
		this.items = append(this.items[0:pos], this.items[pos+1:]...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */