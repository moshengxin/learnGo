package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前系统时间
	now := time.Now()

	// 1. 获取当前时间的 秒级时间戳（从1970-01-01到现在的总秒数）
	secs := now.Unix()

	// 2. 获取当前时间的 纳秒级时间戳（总纳秒数）
	nanos := now.UnixNano()

	// 打印原始时间
	fmt.Println("当前时间:", now)

	// 3. 毫秒级时间戳 = 纳秒 / 1000000
	millis := nanos / 1000000
	fmt.Println("秒时间戳:", secs)
	fmt.Println("毫秒时间戳:", millis)
	fmt.Println("纳秒时间戳:", nanos)

	// 4. 把 秒时间戳 转回时间格式
	fmt.Println("秒戳转时间:", time.Unix(secs, 0))

	// 5. 把 纳秒时间戳 转回时间格式
	fmt.Println("纳秒戳转时间:", time.Unix(0, nanos))
}
