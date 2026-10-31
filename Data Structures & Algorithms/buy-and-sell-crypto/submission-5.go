func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	l, r := 0, 1
	profit := 0

	for r < len(prices) {
		if prices[l] > prices[r] {
			l = r
		}
		curr := prices[r] - prices[l]
		profit = max(curr, profit)
		r++
	}
	return profit
}
