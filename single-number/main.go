package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		seen[nums[i]]++
	}
	for key, value := range seen {
		if value == 1 {
			return key
		}
	}
	return 0
}
