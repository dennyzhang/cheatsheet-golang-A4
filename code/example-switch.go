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
// Updated: Time-stamp: <2018-11-17 11:54:31>
//-------------------------------------------------------------------
// https://tour.golang.org/flowcontrol/9
package main

import (
	"fmt"
	"runtime"
)

func example1()
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// https://gobyexample.com/switch
func example2() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
}
