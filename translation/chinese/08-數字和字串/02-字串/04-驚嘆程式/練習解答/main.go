// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"os"
	"strings"
)

// 注意：你至少應該傳入一個參數

func main() {
	msg := os.Args[1]

	// 重要的是只計算一次
	// 因此，不要呼叫 repeat 函數兩次
	// 呼叫一次就足夠了
	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)

	// 你也可以把這個程式寫得更簡潔
	// 例如：
	//
	// msg := strings.ToUpper(os.Args[1])
	// marks := strings.Repeat("!", len(msg))
	// fmt.Println(marks + msg + marks)
}
