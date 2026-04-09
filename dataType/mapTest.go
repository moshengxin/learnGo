package main

import "fmt"

func main() {

	// map定义
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println("map1:", map1)

	var mp2 = make(map[int]int)
	mp2[1] = 1
	mp2[444] = 12
	mp2[3] = 14

	//
	value, exist := mp2[2]
	val := mp2[1]
	fmt.Println("v1是什么:", mp2, value, exist, val)

	// 删除某个key
	delete(mp2, 1)
	fmt.Println("mp2:", mp2)
}
