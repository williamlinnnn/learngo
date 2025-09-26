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
	"unicode/utf8"
)

func main() {
	// 字串是由位元組（bytes）組成的

	// len 函數會計算字串值中的位元組數
	//
	// 這個字串字面值包含 Unicode 字元
	//
	// 而 Unicode 字元可能佔 1 到 4 個位元組
	// 因此，"İnanç" 的長度是 7 個位元組，而不是 5
	//
	// İ = 2 bytes
	// n = 1 byte
	// a = 1 byte
	// n = 1 byte
	// ç = 2 bytes
	// TOTAL = 7 bytes
	name := "İnanç"
	fmt.Printf("%q is %d bytes\n", name, len(name))

	// 若要取得 UTF-8 編碼字串中的實際字元（或 rune），
	// 你應該這麼做：
	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))
}
