package main

import (
	"fmt"
)

func main() {

	// 变量定义
	var a int = 123
	var b = 3.2
	c := "字符串"
	var d, f = "123", "456"
	fmt.Println(a, b, c, d, f)

	// 常量定义
	const Pi = 3.14
	fmt.Println(Pi)

	// for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for i := 0; i < 2; i++ {
		fmt.Println("循环两次", i)
	}

	for {
		fmt.Println("无条件，就是死循环")
		break
	}

	// if else
	if 1 < 2 {
		fmt.Println("条件判断不用括号，但是一定要{}")
	} else {
		fmt.Println("else分支")
	}

	// switch
	switch i := 2; i {

	case 1:
		fmt.Println("switch 1")
	case 2:
		fmt.Println("switch 2")
	}

	// 循环数组
	for k, v := range []int{1, 2, 3} {
		fmt.Println(k, v)
	}

	arr := [3]string{"111", "222", "333"}
	for i := range arr {
		fmt.Println(arr[i])
	}

	// if语句
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
