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

	for t > 0 {
		t--

		var n, k int
		fmt.Fscanf(reader, "%d %d\n", &n, &k)

		grid := make([][]byte, n)
		for i := 0; i < n; i++ {
			grid[i] = make([]byte, n)
			fmt.Fscanf(reader, "%s\n", &grid[i])
		}

		// 缩减后的网格大小是 n/k x n/k
		reducedSize := n / k
		reducedGrid := make([][]byte, reducedSize)
		for i := 0; i < reducedSize; i++ {
			reducedGrid[i] = make([]byte, reducedSize)
			for j := 0; j < reducedSize; j++ {
				// 取每个块的左上角元素作为缩减后的值
				reducedGrid[i][j] = grid[i*k][j*k]
			}
		}

		// 输出缩减后的网格
		for i := 0; i < reducedSize; i++ {
			fmt.Fprintf(writer, "%s\n", reducedGrid[i])
		}
	}
}
