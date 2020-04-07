package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	previous := 0

	for _, num := range nums {
		previous = int(math.Max(float64(previous+num), float64(num)))
		max = int(math.Max(float64(max), float64(previous)))
	}
	return max
}
