package main

import "fmt"

/*
*
1. 主协程执行到 msg := <-messages 时，channel 是空的
2. 主协程阻塞，等待数据
3. 子 goroutine 执行 messages <- "ping"，发送数据
4. 主协程接收到数据，解除阻塞，继续执行
5. 打印 "ping"
*/
func main() {
	//test1()
	test2()
}

// 无缓冲chan
func test1() {

	// messages的类型是string类型的chan类型
	messages := make(chan string)

	// go开头的，是一个协程，
	go func() { messages <- "ping" }()
	go func() { messages <- "123" }()

	msg := <-messages
	msg2 := <-messages

	// 打印 ping 123或者123 ping，因为两个协程是并发启动执行的，有先后顺序
	fmt.Println(msg, msg2)

}

// 有缓冲chan
func test2() {

	// 最多允许缓存 2 个值,且无需并发接收
	msg := make(chan string, 2)
	msg <- "ping"
	msg <- "xxx"
	fmt.Println(<-msg)
	fmt.Println(<-msg)

}
