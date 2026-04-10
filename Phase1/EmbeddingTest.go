package main

import "fmt"

/**
* 其实是“假继承”，结构体A中，加入结构体B，A就继承了结构体B的所有字段和方法，可以做到代码复用
可以复用多个结构体中的字段和方法，同名字段或者方法，会被覆盖，优先用外层的字段和方法
*/

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
