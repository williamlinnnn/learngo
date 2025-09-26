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
	"strings"
)

// ---------------------------------------------------------
// 練習：改良版驚嘆程式（Improved Banger）
//
//  修改 Banger 程式，使其可以處理 Unicode 字元
//
// 輸入
//  "İNANÇ"
//
// 預期輸出
//  İNANÇ!!!!!
// ---------------------------------------------------------

func main() {
	msg := os.Args[1]

	s := msg + strings.Repeat("!", len(msg))

	fmt.Println(s)
}
