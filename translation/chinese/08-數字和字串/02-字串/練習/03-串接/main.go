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
// 練習：原始字串連接
//
//  1. 透過命令列輸入初始化 name 變數
//
//  2. 使用連接運算子將它與下面的原始字串連接
//
// 注意
//  你應該在正確的位置連接 name 變數
//
// 預期輸出
//  假設你這樣執行程式：
//   go run main.go inanç
//
//  那麼輸出應該如下：
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func main() {
	// 取消註解下面的程式碼
	// name := "從命令列取得名字"

	// 將 `name` 變數替換並連接到下面的 `hi ` 之後

	msg := `hi CONCATENATE-NAME-VARIABLE-HERE!
how are you?`

	fmt.Println(msg)
}
