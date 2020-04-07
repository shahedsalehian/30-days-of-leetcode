package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, num := range nums {
		fmt.Printf("%v ", num)
	}
}
func moveZeroes(nums []int) {
	temp := []int{}
	count := 0
	for _, num := range nums {
		if num != 0 {
			temp = append(temp, num)
			count++
		}
	}
	count = len(nums) - count
	for i := 0; i < count; i++ {
		temp = append(temp, 0)
	}
	for i, val := range temp {
		nums[i] = val
	}
}
