package main

import (
	"errors"
	"fmt"
)

func main() {

	a, err := f1()
	fmt.Println("hello world", a, err)
	//b, err1 := f2()
	//if err1 != nil {
	//	fmt.Println("发生了异常", err1)
	//	return
	//}
	//fmt.Println("hello world", b, err1)

	a3, err3 := f3(42)
	fmt.Println("3333", a3, err3)
}

func f1() (string, error) {
	return "11", nil
}

func f2() (string, error) {
	return "11", errors.New("自定义错误")
}

type argError struct {
	arg  int
	prob string
}

func (arg *argError) Error() string {
	res := "要实现error接口中的Error方法，才能在函数中返回argError错误" + arg.prob
	return res
}

func f3(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
