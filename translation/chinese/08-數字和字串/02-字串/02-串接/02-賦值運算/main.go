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
	name, last := "carl", "sagan"

	// 使用串接
	name += " edward"

	// 等於這個:
	// name = name + " edward"

	fmt.Println(name + " " + last)
}
