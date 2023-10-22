package main

func maxProfit(prices []int) int {
	localMin := 10000

	res := 0
	for _, v := range prices {
		if localMin > v {
			localMin = v
		}
		if margin := v - localMin; margin > res {
			res = margin
		}
	}

	return res
}
