//-------------------------------------------------------------------
// @copyright 2017 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: helloworld.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description :
// https://cheatsheet.dennyzhang.com/cheatsheet-golang-A4
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-10-06 16:40:15>
//-------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3}
	B := &A
	fmt.Println((*B)[2])
	fmt.Println(B)
}
