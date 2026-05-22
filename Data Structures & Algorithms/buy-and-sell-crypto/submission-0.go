func maxProfit(prices []int) int {
	profits := map[int]int{}
	
	for i, v := range prices {
		j := i + 1
		p := 0
		for j < len(prices) {
			if prices[j] - v > p {
				p = prices[j] - v
			}
			j++
		}
		profits[i] = p
	}


	max := 0
	for _, v := range profits {
		if v > max {
			max = v
		}
	}
	return max
}
