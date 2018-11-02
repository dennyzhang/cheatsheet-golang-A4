//-------------------------------------------------------------------
// @copyright 2017 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: tree-trie.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run tree-trie.go
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-11-02 00:37:16>
//-------------------------------------------------------------------
package main

// Blog link: https://code.dennyzhang.com/short-encoding-of-words
// Basic Ideas: Trie (One pass, instead of two)
//     Reverse the words, build trie tree
//     Get the length from the items in the tree
//
//     Note: the expression is not unique
//      Both: "time#bell#" and "bell#time#" are fine
//  Assumptions: no duplicate words
//
// Complexity: Time O(n), Space O(n)
//    n = total count of characters in all words

type TrieNode struct {
   is_leaf bool
   children map[byte]*TrieNode
}

func minimumLengthEncoding(words []string) int {
    // build TrieTree
    res := 0
    leaf_count := 0

    root := TrieNode{children: make(map[byte]*TrieNode)}
    for _, word := range words {
        p := &root
        for i:=len(word)-1; i>=0; i-- {
            ch := word[i]
            q, status := p.children[ch]
            if status == false {
                q = &TrieNode{children: make(map[byte]*TrieNode)}
                p.children[ch] = q
            }
            p = q
            res += 1
            if p.is_leaf == true {
                // revert
                p.is_leaf = false
                leaf_count -= 1
                res -= (len(word) - i)
            }
        }
        p.is_leaf = true
        leaf_count += 1
        if len(p.children) != 0 {
            // revert
            p.is_leaf = false
            leaf_count -= 1
            res -= len(word)
        }
    }
    return res + leaf_count
}
