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
// 練習:球體表面積
//
//  1. 從命令列取得半徑
//  2. 將其轉換為 float64
//  3. 計算球體的表面積
//
// 球體表面積公式
//  area = 4πr²
//  https://en.wikipedia.org/wiki/Sphere#Surface_area
//
// 限制條件
//  使用 `math.Pow` 來計算面積
//  可查看其文件了解使用方式
//  https://golang.org/pkg/math/#Pow
//
// 預期輸出
//  go run main.go 10
//    1256.64
//
//  go run main.go 54.2
//    36915.47
// ---------------------------------------------------------

func main() {
	var radius, area float64

	// 在此添加程式碼
	// ...

	// 別動這裡
	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)
}
