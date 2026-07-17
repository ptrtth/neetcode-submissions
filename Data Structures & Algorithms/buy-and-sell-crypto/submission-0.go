func maxProfit(prices []int) int {
	result := 0

	for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			result = max(result, prices[j] - prices[i])
		}
	}

	return result
}
