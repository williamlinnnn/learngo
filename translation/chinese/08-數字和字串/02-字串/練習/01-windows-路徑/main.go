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
// 練習：Windows 路徑
//
//  1. 修改下列程式
//  2. 改用原始字串（raw string literal）
//
// 提示
//  先執行這個程式看看輸出
//  這樣你就能輕鬆理解它的作用
//
// 預期輸出
//  你的解答應該輸出與這個程式相同的結果
//  不同的是應該使用原始字串
// ---------------------------------------------------------

func main() {
	// 提示：
	// \\ 代表反斜線字元
	// \n 代表換行字元

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println(path)
}
