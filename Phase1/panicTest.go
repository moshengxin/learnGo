package main

import "os" // 用来操作文件

func main() {
	// 1. 直接手动触发 panic，程序立刻崩溃退出
	panic("a problem")

	// 2. 下面这行代码 **永远不会执行**，因为上面已经 panic 了
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err) // 如果创建文件失败，也触发 panic
	}
}
