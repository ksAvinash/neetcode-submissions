func maxProfit(prices []int) int {
	profit, i := 0, 1
	for i < len(prices) {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
		i++
	}
	return profit
}
