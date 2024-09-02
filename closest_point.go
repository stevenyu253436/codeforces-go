package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanf(reader, "%d\n", &n)

		points := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d", &points[j])
		}

		// Skip the remaining line after reading all points.
		fmt.Fscanf(reader, "\n")

		if n == 2 {
			// 當n==2時，檢查兩個點是否相鄰
			if points[1]-points[0] == 1 {
				fmt.Fprintln(writer, "NO")
			} else {
				fmt.Fprintln(writer, "YES")
			}
		} else {
			// 當n>=3時，直接輸出NO
			fmt.Fprintln(writer, "NO")
		}
	}
}
