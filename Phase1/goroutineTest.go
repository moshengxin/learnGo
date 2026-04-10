package main

import (
	"fmt"
	"time"
)

// 普通函数，作为协程执行
func printHello() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("协程执行：%d\n", i)
		time.Sleep(200 * time.Millisecond) // 模拟耗时操作
	}
}

func main() {
	// 启动一个 Go 协程
	go printHello()

	// 主协程继续执行自己的逻辑
	for i := 1; i <= 3; i++ {
		fmt.Printf("主协程执行：%d\n", i)
		time.Sleep(300 * time.Millisecond)
	}

	// 等待子协程执行完毕（简单方案）
	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")
}
