// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// Go 編譯器將這些數字視為整數，
	// 因為整數值中沒有小數部分，
	// 所以結果會變成 1 而不是 1.5

	// 因此，這裡的 ratio 變數是一個 int 型別，
	// 因為 3 除以 2 的結果是一個整數。

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}
