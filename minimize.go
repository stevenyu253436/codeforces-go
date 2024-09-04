package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t) // 讀取測試案例數量

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		// 因為 (c - a) + (b - c) = b - a，因此直接輸出 b - a
		fmt.Println(b - a)
	}
}
