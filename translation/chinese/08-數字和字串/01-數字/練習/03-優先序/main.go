// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// 練習：優先序
//
//  修改運算式，讓它們產生預期的輸出
//
// 限制條件：
//  使用括號來改變優先順序
// ---------------------------------------------------------

func main() {
	// 輸出 20
	fmt.Println(10 + 5 - 5 - 10)

	// 輸出 -16
	fmt.Println(-10 + 0.5 - 1 + 5.5)

	// 輸出 -25
	fmt.Println(5 + 10*2 - 5)

	// 輸出 0.5
	fmt.Println(0.5*2 - 1)

	// 輸出 24
	fmt.Println(3 + 1/2*10 + 4)

	// 輸出 15
	fmt.Println(10 / 2 * 10 % 7)

	// 輸出 40
	// 注意要使用浮點數來解決
	fmt.Println(100 / 5 / 2)
}
