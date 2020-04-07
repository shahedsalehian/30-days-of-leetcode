package main

import (
	"fmt"
	"math"
)

func main() {
	num := 19
	fmt.Println(isHappy(num))
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 {
		m := []int{}
		for n > 0 {
			m = append(m, n%10)
			n = n / 10
		}

		temp := 0
		for i := range m {
			temp += int(math.Pow(float64(m[i]), float64(2)))
		}
		if seen[temp] == true {
			return false
		}
		n = temp
		seen[n] = true
	}

	return true
}
