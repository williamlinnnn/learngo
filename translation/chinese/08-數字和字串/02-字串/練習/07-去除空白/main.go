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
)

// ---------------------------------------------------------
// 練習：去除空白（Trim It）
//
//  1. 查看 strings 套件的文件
//  2. 找一個可以去除給定字串前後空白的函數
//  3. 去除 text 變數的空白並輸出
//
// 預期輸出
//  The weather looks good.
//  I should go and play.
// ---------------------------------------------------------

func main() {
	msg := `
	
	The weather looks good.
I should go and play.



	`

	fmt.Println(msg)
}
