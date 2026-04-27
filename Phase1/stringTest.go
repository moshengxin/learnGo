package main

import (
	"fmt"
	"strconv"
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

	// string转数值类型
	str := "9999"
	i, _ := strconv.ParseInt(str, 10, 8)
	fmt.Println("====", i)

	//直接转
	i2, _ := strconv.Atoi(str)
	fmt.Println("直接====", i2)
}
