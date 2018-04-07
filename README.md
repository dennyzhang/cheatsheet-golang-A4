# awesome-golang-syntax
<a href="https://github.com/DennyZhang?tab=followers"><img align="right" width="200" height="183" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/fork_github.png" /></a>

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com) [![LinkedIn](https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/linkedin.png)](https://www.linkedin.com/in/dennyzhang001) <a href="https://www.dennyzhang.com/slack" target="_blank" rel="nofollow"><img src="http://slack.dennyzhang.com/badge.svg" alt="slack"/></a> [![Github](https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/github.png)](https://github.com/DennyZhang)

File me [tickets](https://github.com/DennyZhang/awesome-golang-syntax/issues) or star [the repo](https://github.com/DennyZhang/awesome-golang-syntax).

See more CheatSheets from Denny: [here](https://github.com/topics/denny-cheatsheets)

Table of Contents
=================

   * [awesome-golang-syntax](#awesome-golang-syntax)
      * [List](#list)
      * [Compact Coding](#compact-coding)
      * [String](#string)
      * [Integer](#integer)
      * [Dict &amp; Set](#dict--set)
      * [Goroutines](#goroutines)
      * [Bit Operator](#bit-operator)
      * [Math](#math)
      * [queue/heapq](#queueheapq)
   * [Code snippets](#code-snippets)
   * [More links](#more-links)

<a href="https://www.dennyzhang.com"><img align="right" width="185" height="37" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>

## Golang Compact Coding

| Name                                | Comment                                     |
| :---------------------------------- | -----------------------------------------   |
| Declare variables with initializers | `var ischecked, v, str  = false, 2, "yes!"` |
  
## List

| Name                        | Comment                                       |
| :----------------------     | --------------------------------------------  |
| Make a list                 | `var a [2]string; a[0]="hello"; a[1]="world"` |
| Create list with fixed size | `T := [6]int{2, 3, 5, 7, 11, 13}`             |
| Create a flexible list      | `T := []string{"a", "c", "b", "d"}`           |
| Sort list                   | `sort.Strings(T); fmt.Print(T)`               |
| Append item                 | `T = append(T, "e")`                          |
| Remove last itme            | `T = T[:len(T)-1]`                            |
| Slices of a list            | `var T2 = T[1:3]` (references of arrays)      |
  
## String

| Name                      | Comment                               |
| :------------------------ | ------------------------------------- |
| String to int             | `i, _ := strconv.Atoi("39038")`       |

## Integer

| Name                      | Comment                        |
| :------------------------ | ------------------------------ |

## Dict & Set

| Name                      | Comment                               |
| :---------------------    | -----------------------------------   |

## Goroutines
| Name                      | Comment                               |
| :---------------------    | -----------------------------------   |

## Bit Operator

| Name                  | Comment             |
| :-------------        | ----------------    |

## File
| Name           | Comment                                         |
| :------------- | ----------------                                |
  
## Math

| Name          | Comment                                   |
| :------------ | -----------------------------------       |

## queue/heapq

| Name                | Comment                                                   |
| :-----------------  | --------------------------------------------------------- |

# Code snippets

# More links
- [A Tour Of Go](https://tour.golang.org/list)
- [The Go Programming Language](https://golang.org/doc/)

- TODO: Need to automatically generate pdf

- License: Code is licensed under [MIT License](https://www.dennyzhang.com/wp-content/mit_license.txt).

<a href="https://www.dennyzhang.com"><img align="right" width="201" height="268" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/denny_201706.png"></a>

<a href="https://www.dennyzhang.com"><img align="right" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>
