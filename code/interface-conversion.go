//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
// https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: interface_conversion.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run interface_conversion.go
// Link: https://www.golang-book.com/books/intro/13
// https://cheatsheet.dennyzhang.com/cheatsheet-golang-A4
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-10-06 16:40:25>
//-------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	// https://stackoverflow.com/questions/26975880/convert-mapinterface-interface-to-mapstringstring
	m := make(map[interface{}]interface{})
	m["foo"] = "bar"

	m2 := make(map[string]string)   

	for key, value := range m {        
		switch key := key.(type) {
		case string:
			switch value := value.(type) {
			case string:
				m2[key] = value
				fmt.Println(key, ":", value)
			}
		}
	}
}
// File: interface_conversion.go ends
