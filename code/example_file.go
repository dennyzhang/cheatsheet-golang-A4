//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
// https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: example_file.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run example_file.go
// Link: https://www.golang-book.com/books/intro/13
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-04-18 23:49:26>
//-------------------------------------------------------------------
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
// File: example_file.go ends
