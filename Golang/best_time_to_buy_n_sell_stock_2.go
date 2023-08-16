package leetcode

// func maxProfit2(prices []int) int {
// 	var lowest, maxProfit = -1, 0

// 	for x, el := range prices {
// 		if el < lowest || lowest == -1 {
// 			lowest = el
// 		}
// 		diff := el - lowest
// 		if lowest < el && el > prices[x + 1] && diff > 0 {
// 			maxProfit += diff
// 			lowest = -1
// 		}
// 	}

// 	return maxProfit
// }

func maxProfit2(prices []int) int {
	var totalProfit = 0

	for i := 0; i < len(prices) - 1; i++ {
		if prices[i+1] > prices[i] {
			totalProfit += prices[i+1] - prices[i]
		}
	}

	return totalProfit
}
