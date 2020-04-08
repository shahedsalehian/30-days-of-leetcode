package main

import (
	"fmt"
)

func main() {
	arr := []int{1,1,2}
	fmt.Println(countElements(arr))
}

func countElements(arr []int) int {
	count := 0
	for _, val := range arr {
		if contains(arr, val + 1) {
			count++
		}
	}	
	return count
} 
 
func contains(arr []int, num int) bool {
	for _, val := range arr {
	   if val == num {
		  return true
	   }
	}
	return false
}