package main

import (
	"fmt"
	"strings"
)

type rect struct {
	width, height int
	color         string
}

// 方法接收者是结构体rect，相当于第一个参数,如果有多个参数，可放在括号里边，相当于第二，第三...第n个参数
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	println(strings.ToUpper("hello world"))

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
