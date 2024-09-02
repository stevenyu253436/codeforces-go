package main

import (
	"fmt"
)

func solve(n int, p []int, s string) []int {
	visited := make([]bool, n)
	F := make([]int, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			// 發現一個新的循環
			cycle := []int{}
			countBlack := 0

			// 遍歷這個循環
			for current := i; !visited[current]; current = p[current] - 1 {
				visited[current] = true
				cycle = append(cycle, current)
				if s[current] == '0' {
					countBlack++
				}
			}

			// 將循環內的所有節點的F(i)設置為循環內黑色節點的數量
			for _, v := range cycle {
				F[v] = countBlack
			}
		}
	}

	return F
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		var n int
		fmt.Scan(&n)

		p := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&p[i])
		}

		var s string
		fmt.Scan(&s)

		F := solve(n, p, s)

		for i := 0; i < n; i++ {
			fmt.Print(F[i], " ")
		}
		fmt.Println()
	}
}
