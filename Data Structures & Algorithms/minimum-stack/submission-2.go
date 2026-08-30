type MinStack struct {
	items []int
	minItem int
}

func Constructor() MinStack {
	return MinStack{
		items: []int{},
		minItem: math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	this.minItem = min(this.minItem, val)
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
	this.minItem = math.MaxInt
	for _, v := range this.items {
		this.minItem = min(this.minItem, v)
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.minItem
}
