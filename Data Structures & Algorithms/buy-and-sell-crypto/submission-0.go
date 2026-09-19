func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	res := 0
	buy := 0
	for buy < len(prices)-1 {
		if prices[buy+1] < prices[buy] {
			buy++
			continue
		}
		sell := buy + 1
		for sell < len(prices) {
			res = max(res, prices[sell]-prices[buy])
			sell++
		}
		buy++
	}
	return res
}
