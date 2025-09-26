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
// 練習:簡化敘述
//
//  簡化程式碼（重構）
//
// 限制條件：
//  只能使用遞增遞減與指派運算
//
// 預期輸出
//  3
// ---------------------------------------------------------

func main() {
	width, height := 10, 2

	width = width + 1
	width = width + height
	width = width - 1
	width = width - height
	width = width * 20
	width = width / 25
	width = width % 5

	fmt.Println(width)
}
