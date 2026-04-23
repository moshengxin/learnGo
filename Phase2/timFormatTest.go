package main

import (
	"fmt"
	"time"
)

func main() {
	// 简写打印函数，方便使用
	p := fmt.Println

	// 获取当前时间
	t := time.Now()

	// 1. 使用标准格式 RFC3339 格式化时间（国际通用标准）
	p("标准格式", t.Format(time.RFC3339))

	// 2. 解析 RFC3339 格式的字符串，转为时间对象
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p("格式化", t1) // 输出解析后的时间

	// 3. 自定义格式输出时间（Go 固定用 2006-01-02 15:04:05 作为参考模板）
	p(t.Format("3:04PM"))                           // 小时:分钟 AM/PM
	p(t.Format("Mon Jan _2 15:04:05 2006"))         // 星期 月 日 时:分:秒 年
	p(t.Format("2006-01-02T15:04:05.999999-07:00")) // 完整时间格式

	// 4. 自定义格式解析时间
	form := "3 04 PM"                    // 定义解析模板
	t2, e := time.Parse(form, "8 41 PM") // 按模板解析字符串
	p(t2)

	// 5. 手动拼接格式化输出时间
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// 6. 格式不匹配，解析失败，输出错误信息
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
