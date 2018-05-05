# cheatsheet-golang-a4
<a href="https://github.com/DennyZhang?tab=followers"><img align="right" width="200" height="183" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/fork_github.png" /></a>

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com) [![LinkedIn](https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/linkedin_icon.png)](https://www.linkedin.com/in/dennyzhang001) <a href="https://www.dennyzhang.com/slack" target="_blank" rel="nofollow"><img src="http://slack.dennyzhang.com/badge.svg" alt="slack"/></a> [![Github](https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/github.png)](https://github.com/DennyZhang)

File me [tickets](https://github.com/DennyZhang/cheatsheet-golang-a4/issues) or star [the repo](https://github.com/DennyZhang/cheatsheet-golang-a4).

See more CheatSheets from Denny: [#denny-cheatsheets](https://github.com/topics/denny-cheatsheets)

Table of Contents
=================

   * [cheatsheet-golang-a4](#cheatsheet-golang-a4)
      * [Syntax Sugar: From Python To Golang](#syntax-sugar-from-python-to-golang)
      * [Golang Compact Coding](#golang-compact-coding)
      * [Array/List/Slice](#arraylistslice)
      * [String](#string)
      * [Integer/Float](#integerfloat)
      * [Conversion](#conversion)
      * [Dict/Hashmap/Map](#dicthashmapmap)
      * [Goroutines](#goroutines)
      * [Files &amp; Folders](#files--folders)
      * [Bit Operator](#bit-operator)
      * [Math](#math)
      * [queue/heapq](#queueheapq)
   * [Code snippets](#code-snippets)
   * [More links](#more-links)

<a href="https://www.dennyzhang.com"><img align="right" width="185" height="37" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>

**Golang CheatSheet**: https://github.com/DennyZhang/cheatsheet-golang-A4

## Syntax Sugar: From Python To Golang
| Name           | Python                | Golang                                             |
| :------------- | --------------------- | -------------------------------------------------- |
| sum slice      | `sum([1, 2, 3])`      | `sum := 0; for i := range nums { sum += nums[i] }` |
| Get last item  | `nums[-1]`            | `nums[len(nums)-1]`                                |
| For            | `for i in range(10):` | `for i := 0; i < 10; i++`                          |
| Loop list      | `for num in [1, 2]`   | `for num := range[]int{1, 2} { fmt.Print(num) }`   |
| Loop string    | `for ch in str:`      | `for _, ch := range str { fmt.Print(ch) }`         |
| Iterator       | `for num in nums:`    | `for _, num := range nums {fmt.Print(num)}`        |
| While          | `while isOK:`         | `for isOK`                                         |
| Reverse list   | `nums[::-1]`          | Need to create your own function. Weird!           |
| Get min        | `min(2, 6, 5)`        |                                                    |

## Golang Compact Coding

| Name                                | Comment                                     |
| :---------------------------------- | ------------------------------------------- |
| Declare variables with initializers | `var ischecked, v, str  = false, 2, "yes!"` |
| One line if statement               | `if a >= 1 { fmt.Print("yes") }`            |

## Array/List/Slice

| Name                            | Comment                                       |
| :-----------------------------  | --------------------------------------------  |
| Make a array                    | `var a [2]string; a[0]="hello"; a[1]="world"` |
| Create array with given values  | `T := [6]int{2, 3, 7, 5, 11, 13}`             |
| Create array with given values  | `T := []string{"a", "c", "b", "d"}`           |
| Create dynamically-sized arrays | `a := make([]int, 5)`                         |
| Create dynamically-sized arrays | `a := make([]int, 1, 5)` // 5 is capacity     |
| Sort string array               | `sort.Strings(T); fmt.Print(T)`               |
| Sort int array                  | `sort.Ints(l[:])`                             |
| Append item                     | `T = append(T, "e")`                          |
| Append items                    | `T = append(T, "e", "b", "c")`                |
| Remove last item                | `T = T[:len(T)-1]`                            |
| Slices of a array               | `var T2 = T[1:3]` // Notice: it's a reference |
| Copy a list                     | `b := make([]T, len(a)); copy(b, a)`          |
  
## String

[Package strings](https://golang.org/pkg/strings/)

| Name                         | Comment                                                         |
| :------------------------    | --------------------------------------------------------------- |
| Format string                | `fmt.Sprintf("at %v, %s", e.When, e.What)`                      |
| Format string                | `fmt.Printf("int: %d, float: %f, bool: %t\n", 123, 78.9, true)` |
| Split string                 | `var L = strings.Split("hi,golang", ",")`                       |
| Replace string               | `var str2 = strings.Replace("hi,all", ",", ";", -1)`            |
| Replace string               | `strings.Replace("aaaa", "a", "b", 2)` //bbaa                   |
| Split string by separator    | `strings.Split(path, " ")`                                      |
| Count characters             | `strings.Count("test", "t")`                                    |
| Substring                    | `strings.Index("test", "e")`                                    |
| Join string                  | `strings.Join([]string{"a","b"}, "-")`                          |
| Repeat string                | `strings.Repeat("a", 2)` // aa                                  |
| Tolower                      | `strings.ToLower("TEST")`                                       |
| Trim whitespace in two sides | `strings.TrimSpace("\t Hello world!\n ")`                       |
| Trim trailing whitespace     | `strings.TrimRight("\t Hello world!\n ", "\n ")`                |

## Integer/Float

| Name                | Comment                                                             |
| :------------------ | ------------------------------------------------------------------- |
| Int max             | `MaxInt32  = 1<<31 - 1` [golang math](https://golang.org/pkg/math/) |
| Int min             | `MinInt32 = -1 << 31`   [golang math](https://golang.org/pkg/math/) |

## Conversion
| Name                    | Comment                                       |
| :---------------------- | --------------------------------------------- |
| Convert string to int   | `i, _ := strconv.ParseInt("12345", 10, 64)`   |
| Convert string to int   | `i, err := strconv.Atoi("-42")`               |
| Convert string to float | `f, _ := strconv.ParseFloat("3.1415", 64)`    |
| Convert int to float    | `0.5*float32(age)+7>= float32(age2)`          |
| Convert int to string   | `s := strconv.Itoa(-42)`                      |

## Dict/Hashmap/Map

| Name                  | Comment                          |
| :-------------------  | -------------------------------- |
| Create dict           | `map[string]int{"a": 1, "b": 2}` |
| Create dict           | `make(map[string]int)`           |
| Check existence       | `_, ok := m[k]`                  |
| Delete key            | `delete(m, "k1")`                |
| Create a map of lists | `m := make(map[string][]string)` |
  
## Goroutines
| Name                   | Comment                                      |
| :--------------------- | -------------------------------------------- |
| Basic goroutine        | [example_goroutine.go](example_goroutine.go) |

## Files & Folders
| Name                   | Comment                            |
| :--------------------- | ---------------------------------- |
| Read/Write files       | [example_file.go](example_file.go) |

## Bit Operator

| Name           | Comment                       |
| :------------- | ----------------------------- |
| Shift left     | `fmt.Print(1 << 10)` // 1024  |
| Shift right    | `fmt.Print(1024 >> 3)` // 128 |
  
## Math

| Name          | Comment                  |
| :------------ | ------------------------ |

## queue/heapq

| Name            | Comment                |
| :-------------- | ---------------------- |

# Code snippets
- Create 2D arrays
```
// static
board := [][]string{
         []string{"_", "_", "_"},
         []string{"_", "_", "_"},
         []string{"_", "_", "_"},
}

// dynamic
a := make([][]uint8, dy)
for i := range a {
    a[i] = make([]uint8, dx)
}
```

- Logging
```
import "github.com/op/go-logging"
log := logging.MustGetLogger("my-app")
log.Info("Some info...")
log.Warning("Some warning...")
log.Error("Some error!")
log.Critical("Some critical!")
```

- struct
```
type Point struct {
	X, Y int
}

var (
	v1 = Point{10, 8}
	v2 = Point{X: 1}  // Y would be 0
	v3 = Point{}      // Both X and Y is 0
	p  = &Point{10, 8} // reference: type *Point
)

func main() {
	fmt.Println(p, v1, v2, v3)
}
```

- Goroutines & Channels
```
// Goroutines
go func() {
  // do something
}
```

```
// Channels
c := make(chan T [, capacity ])
c <- t // blocks on unbuffered channels until another routine receives the value

d := <-c // blocks on unbuffered channels until another routine sends the value

close(c)
```

# More links
- [A Tour Of Go](https://tour.golang.org/list)
- [The Go Programming Language](https://golang.org/doc/)
- https://github.com/a8m/go-lang-cheat-sheet

TODO: Need to automatically generate A4 pdf

License: Code is licensed under [MIT License](https://www.dennyzhang.com/wp-content/mit_license.txt).

<a href="https://www.dennyzhang.com"><img align="right" width="201" height="268" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/denny_201706.png"></a>

<a href="https://www.dennyzhang.com"><img align="right" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>
