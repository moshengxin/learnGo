package main

import (
	"fmt"  // 输入输出打印包
	"time" // Go语言时间处理标准库
)

func main() {
	// 简写 fmt.Println，后续用 p 代替，少打字
	p := fmt.Println

	// 1. 获取当前本地时间
	now := time.Now()
	p("当前系统时间：", now)

	// 2. 手动构造一个指定时间
	// 参数：年, 月, 日, 时, 分, 秒, 纳秒, 时区
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("自定义指定时间：", then)

	// 3. 拆解时间的各个字段
	p("年：", then.Year())
	p("月：", int(then.Month())) // 输出的是Month枚举，对应英文单词，可以用int()转换成数字
	p("日：", then.Day())
	p("时：", then.Hour())
	p("分：", then.Minute())
	p("秒：", then.Second())
	p("纳秒：", then.Nanosecond())
	p("时区：", then.Location())

	// 获取当天是星期几
	p("星期：", then.Weekday())
	p("当前是星期", int(now.Weekday()))

	// 4. 时间先后比较
	p("then 是否在 now 之前？", then.Before(now)) // 早于 → true
	p("then 是否在 now 之后？", then.After(now))  // 晚于 → false
	p("两个时间是否相等？", then.Equal(now))         // 相等 → false

	// 5. 计算两个时间的差值 now - then
	diff := now.Sub(then)
	p("两个时间相差：", diff)

	// 把时间差转换成不同单位
	p("相差小时数：", diff.Hours())
	p("相差分钟数：", diff.Minutes())
	p("相差秒数：", diff.Seconds())
	p("相差纳秒数：", diff.Nanoseconds())

	// 6. 时间加减运算
	p("当前时间加指定时间：", time.Now().Add(2*time.Hour))
}
