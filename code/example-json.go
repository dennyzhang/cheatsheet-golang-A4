//-------------------------------------------------------------------
// @copyright 2018 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: example_json.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run example_json.go
// Link: https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
// https://cheatsheet.dennyzhang.com/cheatsheet-golang-A4
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-10-06 16:40:26>
//-------------------------------------------------------------------
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Page struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Url   string `json:"url"`
}

func (p Page) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func main() {

    pages := getPages()
    for _, p := range pages {
        fmt.Println(p.toString())
    }

    fmt.Println(toJson(pages))
}

func getPages() []Page {
    raw, err := ioutil.ReadFile("./misc/pages.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}
// File: example_json.go ends
