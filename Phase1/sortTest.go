package main

import (
	"fmt"
	"sort" // Go 内置排序包，不用自己写排序算法
)

func main() {
	// 1. 定义一个字符串切片
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)            // 对字符串切片进行**升序排序**
	fmt.Println("Strings:", strs) // 输出：[a b c]

	// 2. 定义一个整数切片
	ints := []int{7, 2, 4}
	sort.Ints(ints)               // 对整数切片进行**升序排序**
	fmt.Println("Ints:   ", ints) // 输出：[2 4 7]

	// 3. 判断一个切片是否已经排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s) // 输出：true
}
