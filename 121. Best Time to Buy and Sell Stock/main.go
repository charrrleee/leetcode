package _21__Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if minPrice > price {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit

}
