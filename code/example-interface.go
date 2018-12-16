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
// Updated: Time-stamp: <2018-12-16 11:07:38>
//-------------------------------------------------------------------
package main

import (
	"fmt"
)

type Car interface {
   Drive() string
}

type Toyota struct {
	name string
}

type Honda struct {
	name string
}

func (p Toyota) Drive() string {
   return "110mph";
}

func (p Honda) Drive() string {
   return "120mph";
}

func main() {
	t := Toyota{name: "toyota"}
	fmt.Println(t.name+" "+t.Drive())

	h := Honda{name: "honda"}
	fmt.Println(h.name+" "+h.Drive())
}
