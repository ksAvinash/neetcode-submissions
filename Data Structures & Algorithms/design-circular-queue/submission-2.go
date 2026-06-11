type MyCircularQueue struct {
	front int
	rear int
	items []int
	max int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
		front: -1,
		rear: -1,
		max: k,
		items: make([]int, k, k),
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
    if this.IsEmpty() {
		this.rear = 0
		this.front = 0
	} else {
		this.rear = (this.rear + 1) % this.max
	}
	this.items[this.rear] = value
	// fmt.Println("EnQueue", this)
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.rear { // single element
		this.front = -1
		this.rear = -1
		this.items = make([]int, this.max, this.max)
	} else {
		this.front = (this.front + 1)%this.max
	}
    // fmt.Println("DeQueue", this)
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
    return this.items[this.front]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
    return this.items[this.rear]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.front == -1 && this.rear == -1
}


func (this *MyCircularQueue) IsFull() bool {
	if this.IsEmpty() {
		return false
	}
    return (this.rear + 1)%this.max == this.front
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 