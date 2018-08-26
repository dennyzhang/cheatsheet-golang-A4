//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: example_goroutine.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run example_goroutine.go
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-04-07 16:15:49>
//-------------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
// File: example_goroutine.go ends
