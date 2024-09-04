package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t) // 讀取測試用例的數量

	for i := 0; i < t; i++ {
		var l, r, L, R int
		fmt.Scan(&l, &r) // Alice 的區間
		fmt.Scan(&L, &R) // Bob 的區間

		// 如果區間不重疊，最少鎖一扇門
		if r < L || R < l {
			fmt.Println(1)
		} else {
			// 區間有重疊，計算需要鎖門的數量
			// fmt.Println(int(math.Max(float64(r+1-L), float64(R+1-l))) + 1)
			if l == L && r == R {
				fmt.Println(r - l)
			} else if l >= L && r <= R {
				// [l, r] 完全包含在 [L, R] 中
				if r == R {
					fmt.Println(r - l + 1)
				} else if l == L {
					fmt.Println(r - l + 1)
				} else {
					fmt.Println(r - l + 2)
				}
			} else if L >= l && R <= r {
				if r == R {
					fmt.Println(R - L + 1)
				} else if l == L {
					fmt.Println(R - L + 1)
				} else {
					fmt.Println(R - L + 2)
				}
			} else if L >= l && r < R {
				fmt.Println(r - L + 2)
			} else if l >= L && R < r {
				fmt.Println(R - l + 2)
			}
		}
	}
}
