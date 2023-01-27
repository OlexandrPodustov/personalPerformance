package main

import "fmt"

func main() {
}

func maxProfit(prices []int) int {
	totalProfit := 0
	min, max := prices[0], prices[len(prices)-1]
	actions := []int{}
	fmt.Println("")
	defer fmt.Println("")

	fmt.Println("min, max", min, max)
	for i := len(prices) - 1; i > 1; i-- {
		fmt.Println("i, prices[i], prices[i-1]", i, "aa", prices[i], prices[i-1])

		if prices[i] > prices[i-1] {
			fmt.Println("skip")

			continue
		}
		if len(actions)%2 == 1 {
			// we can buy
		} else {
			// we can sell
		}

	}

	fmt.Println("min, max", min, max)
	if len(actions)%2 == 1 {
		actions = actions[1:]
	}
	for _, v := range actions {
		totalProfit += v
	}
	fmt.Println("totalProfit", totalProfit)

	return totalProfit
}
