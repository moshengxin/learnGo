package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", strings.Contains("test", "es"))
	println(strings.ToUpper("test"))
	println(strings.Compare("x", "b"))
	println(strings.ContainsAny("test", "x"))
	println(strings.Count("test", "t"))
	println(strings.Cut("test", "a"))
	fmt.Println(strings.Split("1,6,4,g", ",")[0])
	println(strings.ReplaceAll("test", "t", "3"))
	println(strings.Replace("test", "t", "3", 1))
}
