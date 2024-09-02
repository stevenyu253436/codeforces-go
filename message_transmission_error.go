package main

import (
	"fmt"
	"strings"
)

func findOverlap(t string) (string, bool) {
	n := len(t)
	// fmt.Println("Input string:", t) // 打印輸入字串
	for i := n - 1; i > n/2; i-- {
		substr := t[len(t)-i:]
		// fmt.Printf("Checking overlap with length %d: '%s'\n", i, substr) // 打印當前正在檢查的重疊部分
		if strings.HasPrefix(t, substr) {
			// fmt.Println("Found overlap with:", substr) // 打印找到的重疊部分
			return t[len(t)-i:], true
		}
	}
	return "", false
}

func main() {
	var t string
	fmt.Scan(&t)

	s, found := findOverlap(t)
	if found {
		fmt.Println("YES")
		fmt.Println(s)
	} else {
		fmt.Println("NO")
	}
}
