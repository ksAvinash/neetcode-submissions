type MinStack struct {
	items []int
	min *int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.min == nil || val < *this.min {
		this.min = &val
	}
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.items = this.items[0:this.Len()-1]
	if this.min == nil || top == *this.min {
		if this.Len() == 0 {
			this.min = nil
		} else {
			this.min = &this.items[0]
			for _, v := range this.items {
				if v < *this.min {
					this.min = &v
				}
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.items[this.Len()-1]
}

func (this *MinStack) Len() int {
	return len(this.items)
}

func (this *MinStack) GetMin() int {
	return *this.min
}
