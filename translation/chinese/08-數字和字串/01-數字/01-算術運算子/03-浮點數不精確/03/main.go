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
	// 當你在計算中同時使用浮點數和整數時
	// 結果一定會變成浮點數

	ratio := 3.0 / 2

	// OR:
	// ratio = 3 / 2.0

	// OR - if 3 is inside an int variable:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}
