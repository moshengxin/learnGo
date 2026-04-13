package main

import "fmt"

// 记忆点：chan<-  只发不收

// 1. 这个函数：pings 通道 **只能发** (chan<- string)
func ping(pings chan<- string, msg string) {
	pings <- msg // 往通道发消息
}

// 2. 这个函数：pings 只能收，pongs 只能发
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // 从 pings 收消息
	pongs <- msg   // 发到 pongs 通道
}

func main() {
	// 创建两个带缓冲的通道（缓冲=1）
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message") // 第一步：发
	pong(pings, pongs)            // 第二步：转发
	fmt.Println(<-pongs)          // 第三步：收并打印
}
