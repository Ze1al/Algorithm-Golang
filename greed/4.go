package greed

func maxProfit(prices []int) int {
	var res int
	for i := 1; i < len(prices); i++ {
		res += max(prices[i]-prices[i-1], 0)
	}

	return res
}
