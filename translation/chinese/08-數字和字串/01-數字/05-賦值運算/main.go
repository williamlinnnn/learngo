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

func main() {
	width, height := 5., 12.

	// 計算長方形面積
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // 面積減 10
	area = area + 10 // 面積加 10
	area = area * 2  // 面積成 2
	area = area / 2  // 面積除以 2
	fmt.Printf("area=%g\n", area)

	// // 賦值運算
	area -= 10 // 面積減 10
	area += 10 // 面積加 10
	area *= 2  // 面積成 2
	area /= 2  // 面積除以 2
	fmt.Printf("area=%g\n", area)

	// 找出餘數
	// 因為面積是浮點數，這無法運作
	// area %= 7

	// 這個可以運作
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}
