package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. 创建文件
	f := createFile("defer.txt")

	// 🔥 核心：defer 延迟执行
	// 意思：等 main 函数结束前,无论程序是否出错，都会执行，**自动执行 closeFile(f)**
	defer closeFile(f)

	// 2. 写入数据
	writeFile(f)
}

// 创建文件
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// 写入文件
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "123")
}

// 关闭文件
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
