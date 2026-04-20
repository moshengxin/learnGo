package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个工作函数：模拟任务执行
func workerTest(id int) {
	fmt.Printf("workerTest %d starting\n", id) // 开始执行

	time.Sleep(time.Second) // 模拟耗时 1 秒

	fmt.Printf("workerTest %d done\n", id) // 执行完成
}

func main() {
	var wg sync.WaitGroup // 创建一个等待组：用来等所有协程结束

	// 循环 1~5，启动 5 个协程
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 每启动一个协程，计数器 +1

		i := i // 关键：给每个循环创建独立的 i 变量（避免循环变量复用问题）

		// 启动一个 goroutine（轻量级线程）
		go func() {
			defer wg.Done() // 协程结束时，计数器 -1
			workerTest(i)   // 执行工作函数
		}()
	}

	wg.Wait() // 阻塞主线程，直到所有协程的 Done() 把计数器减到 0
}
