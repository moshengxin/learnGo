package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	res := add(1, 3)
	fmt.Println("res:", res)

	res = sub(1, 3, "")
	fmt.Println("res:", res)

	u := User{name: "zm", age: 111}
	test(u)
	fmt.Printf("值传递，方法结束后user实体不会改变: %+v\n", u)

	u2 := User{name: "zm", age: 111}
	fmt.Println("调用方法前: %+v\n", u2)
	testYinYong(&u2)
	fmt.Printf("引用传递，方法结束后user实体会改变: %+v\n", u2)

	add, sub := addAndSub(5, 2)
	fmt.Println("add:", add)
	fmt.Println("sub:", sub)

	fmt.Println(`和是`, sum(1, 2, 97))
	arr := []int{22, 33, 0, 0}
	fmt.Println(`和是`, sum(arr...))

	// 递归：一直调用自己，方法里边有return，防止死循环
	print(fact(7))

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func add(a int, b int) int {
	return a + b
}

// 多个连续的，可以只在最后指定参数的类型
func sub(a, b int, c string) int {
	fmt.Println("c:", c)
	return a - b
}

// 值传递，方法内部改动u，不影响原来的u
func test(u User) {
	nameStr := u.name
	fmt.Println("nameStr", nameStr)

}

// 引用传递，方法内部改动u，会对原来的u，有影响
func testYinYong(u *User) {
	u.name = "test"
	u.age = 18
	fmt.Printf("打印user实体: %+v\n", u)
}

// 多返回值
func addAndSub(a, b int) (int, int) {
	return a + b, a - b
}

// 可变参数，接收一个切片集
func sum(parameters ...int) int {
	sum := 0
	for _, v := range parameters {
		sum += v
	}
	return sum
}
