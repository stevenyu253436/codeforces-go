package main

import (
	"fmt"
)

func alternatingSum(arr []int) int {
	sum := 0
	for i, num := range arr {
		if i%2 == 0 {
			sum += num
		} else {
			sum -= num
		}
	}
	return sum
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		result := alternatingSum(arr)
		fmt.Println(result)
	}
}
