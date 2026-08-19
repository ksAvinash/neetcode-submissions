func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	l, r := 0, 1 // left=buy right=sell
	maxProfit := 0

	for r < len(prices) {
		profit := prices[r] - prices[l]
		maxProfit = max(maxProfit, profit)
		
		if profit < 0 {
			l = r
		}
		r++
	}
	return maxProfit
}
