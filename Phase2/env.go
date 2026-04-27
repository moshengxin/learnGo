package main

import (
	"fmt"     // 打印输出
	"os"      // 操作系统相关库，专门操作环境变量
	"strings" // 字符串分割工具
)

func main() {
	// 1. 设置环境变量：key = FOO，value = 1
	os.Setenv("FOO", "1")

	// 2. 获取指定环境变量的值
	fmt.Println("FOO:", os.Getenv("FOO")) // 输出已设置的 FOO
	fmt.Println("BAR:", os.Getenv("BAR")) // 输出未设置的 BAR，结果为空

	fmt.Println() // 换行

	// 3. 获取系统所有环境变量（返回 []string 切片）
	// 格式：key=value
	for _, e := range os.Environ() {
		// 按 = 分割字符串，只分割成 2 段（防止 value 里有 =）
		pair := strings.SplitN(e, "=", 2)
		// 只打印环境变量的 名称(key)
		fmt.Println(pair[0])
	}

	for i, s := range os.Environ() {
		fmt.Println("key==", i, "value==", s)
	}
}
