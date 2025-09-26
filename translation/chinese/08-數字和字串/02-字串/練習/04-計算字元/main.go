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
	"os"
)

// ---------------------------------------------------------
// 練習：計算字元數
//
//  1. 修改下列程式以支援 Unicode 字元
//
// 輸入
//  "İNANÇ"
//
// 預期輸出
//  5
// ---------------------------------------------------------

func main() {
	// 目前它回傳 7
	// 因為它計算的是位元組...
	// 應該改為計算 rune（碼位）
	//
	// 當你用 "İNANÇ" 執行時，它應該回傳 5，而不是 7

	length := len(os.Args[1])
	fmt.Println(length)
}
