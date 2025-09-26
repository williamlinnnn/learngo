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

// 注意：你至少應該傳入一個參數

func main() {
	msg := os.Args[1]
	l := len(msg)

	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)

	fmt.Println(s)
}
