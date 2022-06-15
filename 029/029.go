package main

import "fmt"

func main() {
	prices := []int{5, 2, 4, 0, 1}
	k := 2
	fmt.Println(getMaxProfit(prices, k))
}

func getMaxProfit(prices []int, k int) int {
	if len(prices) <= 1 || k <= 1 {
		return 0
	}

	// We can trade all profits so get max profits
	if k >= len(prices) {
		return getOnlyProfits(prices)
	}

	// Else we need to be smarter
	local := make([]int, k+1)
	global := make([]int, k+1)
	for i := 1; i < len(prices); i++ {
		changeToPrevDay := prices[i] - prices[i-1]
		for j := k; j >= 1; j-- {
			local[j] = max(global[j-1]+max(changeToPrevDay, 0), local[j]+changeToPrevDay)
			global[j] = max(local[j], global[j])
		}
	}
	return global[k]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func getOnlyProfits(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		changeToPrevDay := prices[i] - prices[i-1]
		if changeToPrevDay > 0 {
			res += changeToPrevDay
		}
	}
	return res
}
