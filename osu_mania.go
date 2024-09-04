package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用帶緩存的讀取和寫入來提高效率
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t) // 讀取測試案例數量

	for t > 0 {
		t--

		var n int
		fmt.Fscanf(reader, "%d\n", &n) // 讀取行數

		// 使用一個切片來儲存每行的結果
		result := make([]int, n)

		for i := 0; i < n; i++ {
			var row string
			fmt.Fscanf(reader, "%s\n", &row) // 讀取每一行

			// 找到 '#' 的列號，並將其存儲在對應的結果中
			for j := 0; j < 4; j++ {
				if row[j] == '#' {
					result[n-i-1] = j + 1 // 因為列號從 1 開始，所以 +1
					break
				}
			}
		}

		// 輸出結果
		for i := 0; i < n; i++ {
			fmt.Fprintf(writer, "%d ", result[i])
		}
		fmt.Fprintln(writer) // 每個測試案例輸出後換行
	}
}
