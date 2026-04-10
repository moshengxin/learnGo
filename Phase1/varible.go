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

	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// 数组 不赋值。默认是0值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	fmt.Println(arr)

	// 数组长度
	fmt.Println(len(arr))

	// 切片 初始化一个固定长度
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("emp:", s)

	// 数组和切片区别
	array := [5]string{}
	slice := []string{}          // 长度为0
	slice = append(slice, "123") // append后，长度为1
	slice[0] = "223"
	slice[10] = "999"
	fmt.Println(len(array), len(slice))

}
