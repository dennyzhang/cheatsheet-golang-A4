//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// Author : Denny <https://www.dennyzhang.com/contact>
// Description :
// https://cheatsheet.dennyzhang.com/cheatsheet-golang-A4
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2019-10-09 10:00:28>
//-------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	// use fixed size array as key, or define a type of struct
	m := map[[2]int]bool{}
	m[[2]int{2, 2}] = true
	fmt.Println("Hello, playground", m)
}
