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
// 練習:圓面積
//
//  根據給定的半徑計算圓的面積
//
// 圓面積公式
//  area = πr²
//  https://en.wikipedia.org/wiki/Area#Circles
//
// 提示
//  可以使用 `math.Pi` 作為 PI
//
// 預期輸出
//  半徑: 10 -> 面積: 314.1592653589793
//
// 額外練習！
//  1. 將面積輸出為 314.16
//  2. 要做到這一點，你需要使用正確的 Printf 格式化字元 :)
//      而不是下面的 `%g`
//
//    預期輸出
//     radius: 10 -> area: 314.16
// ---------------------------------------------------------

func main() {
	var (
		radius = 10.
		area   float64
	)

	// ADD YOUR CODE HERE
	// ...

	// DO NOT TOUCH THIS
	fmt.Printf("radius: %g -> area: %g\n", radius, area)
}
