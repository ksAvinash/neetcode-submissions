type StockSpanner struct {
	items []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		items: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {

	less, j := 1, len(this.items)-1
	for j >= 0 {
		if this.items[j] <= price {
			less++
		} else {
			break
		}
		j--
	}
	this.items = append(this.items, price)
	return less
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor()
 * param1 := obj.Next(price)
 */
 