// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(
		"hello" + ", " + "how" + " " + "are" + " " + "today?",
	)

	// 可以結合字串及原始字串
	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are` + ` ` + "today?",
	)

	// ------------------------------------------
	// 非字串轉為字串
	// ------------------------------------------

	eq := "1 + 2 = "
	sum := 1 + 2

	// 無效的操作
	// 字串連接運算子只能用於字串
	// fmt.Println(eq + sum)

	// 你需要使用 strconv.Itoa 進行轉換
	// Itoa = 整數轉 ASCII

	fmt.Println(eq + strconv.Itoa(sum))

	//

	// 錯誤操作
	// eq = true + " " + false

	eq = strconv.FormatBool(true) +
		" " +
		strconv.FormatBool(false)

	fmt.Println(eq)
}
