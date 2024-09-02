package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// 三個兄弟的編號總和為 6，減去 a 和 b 的值即為缺席的兄弟編號
	lateBrother := 6 - (a + b)

	fmt.Println(lateBrother)
}
