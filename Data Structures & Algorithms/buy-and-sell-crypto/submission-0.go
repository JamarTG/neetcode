func maxProfit(prices []int) int {

		l := 0
		r := 1

		maxProfit := 0

		for r <= len(prices) - 1 {
			profit := prices[r] - prices[l]
			maxProfit = max(maxProfit,profit)

			if profit < 0 {
				l = r
			} 

			r += 1
		}

		return maxProfit
}
