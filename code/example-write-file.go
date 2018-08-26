//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
// https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: example_write_file.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run example_write_file.go
// Link: https://www.golang-book.com/books/intro/13
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-07-10 11:58:04>
//-------------------------------------------------------------------
package main

import (
	"fmt"
	"os"
)

func main() {
	// https://stackoverflow.com/questions/7151261/append-to-a-file-in-go
	file, err := os.OpenFile("/tmp/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// handle the error here
		fmt.Print("Cannot create file", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "hello: %s\n", "world")
	fmt.Fprint(file, "closing the file\n")
}
// File: example_write_file.go ends
