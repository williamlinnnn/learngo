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
	// 字串與原始字串（raw string literal）的類型相同
	// 它們都是 string
	//
	// 因此，它們都可以作為字串值使用
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	// 字串
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	// 原始字串
	s = `
<html>
	<body>"Hello"</body>
</html>`

	fmt.Println(s)

	// windows 路徑
	fmt.Println("c:\\my\\dir\\file") // 字串
	fmt.Println(`c:\my\dir\file`)    // 原始字串
}
