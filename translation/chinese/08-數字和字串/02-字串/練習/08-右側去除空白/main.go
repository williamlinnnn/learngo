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
// 練習：右側去除空白（Right Trim It）
//
//  1. 查看 strings 套件的文件
//  2. 找一個可以只去除字串最右側空白的函數
//  3. 只去除右側的空白
//  4. 輸出它包含的字元數
//
// 限制條件
//  你的程式應該能處理 Unicode 字串值
//
// 預期輸出
//  5
// ---------------------------------------------------------

func main() {
	// 目前它輸出 17
	// 應該輸出 5

	name := "inanç           "
	fmt.Println(len(name))
}
