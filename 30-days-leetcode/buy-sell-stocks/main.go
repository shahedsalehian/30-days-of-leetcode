package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	total := 0
	for i, num := range prices {
		if i == 0 {
			continue
		}
		profit := num - prices[i-1]
		if profit > 0 {
			total += profit
		}
	}
	return total
}
