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
// 練習：列印 JSON
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
	// \t 代表 TAB 字元
	// \n 代表換行字元
	// \" 代表雙引號字元

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	fmt.Println(json)
}
