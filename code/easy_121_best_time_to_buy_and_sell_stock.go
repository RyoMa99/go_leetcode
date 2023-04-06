package code

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	max_profit := make([]int, 0)
	max_profit = append(max_profit, 0)
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - min; diff > max_profit[i-1] {
			max_profit = append(max_profit, diff)
		} else {
			max_profit = append(max_profit, max_profit[i-1])
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	return max_profit[len(prices)-1]
}

// @lc code=end
