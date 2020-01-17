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
// Updated: Time-stamp: <2020-01-16 16:55:58>
//-------------------------------------------------------------------
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
	fmt.Println("Starting the application...")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://myserver.abc.com:9000/v1/clusters", nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbXXXX")
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application...")
}
