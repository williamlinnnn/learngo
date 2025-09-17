// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// 更多教學  : https://learngoprogramming.com
// 親自訓練  : https://www.linkedin.com/in/inancgumus/
// 推特追蹤: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// 當整數 (integer) 和 浮點數 (float) 在同一個運算式中一起使用時
	// 結果一定會變成浮點數 (float)。
	fmt.Println(8 * -4.0) // -32.0 not -32

	// 兩個整數相除會得到整數結果
	fmt.Println(-4 / 2)

	// 餘數運算子
	// 只能用於整數
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) //錯誤

	// 加法與減法運算子
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// 負號運算子（取相反數）
	fmt.Println(-(-2))
	fmt.Println(- -2) // 這樣寫也可以
}
