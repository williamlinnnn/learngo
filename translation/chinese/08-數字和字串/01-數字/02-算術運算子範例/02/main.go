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
	// ratio 的值是多少？
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	// 解釋
	// 上面的運算式等同於：
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	// 如何解決?
	//
	// 記住，只要其中一個值是浮點數
	// 結果就會是浮點數
	ratio = float64(3) / 2
	fmt.Println(ratio)

	// or
	ratio = 3.0 / 2
	fmt.Println(ratio)
}
