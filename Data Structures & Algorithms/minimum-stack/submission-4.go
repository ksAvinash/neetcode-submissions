type MinStack struct {
	items []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		items: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.items = this.items[:len(this.items)-1]

	if top == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
