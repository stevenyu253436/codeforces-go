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

		var n int
		fmt.Fscanf(reader, "%d\n", &n)

		// 計算最少的動物數量
		minAnimals := n / 4
		if n%4 != 0 {
			minAnimals = minAnimals + 1
		}

		fmt.Fprintln(writer, minAnimals)
	}
}
