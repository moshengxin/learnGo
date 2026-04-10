package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 1. 获取当前时间 ===")
	now := time.Now()
	fmt.Println("当前时间:", now)
	fmt.Println("格式化:", now.Format("2006-01-02 15:04:05"))
	fmt.Println()

	fmt.Println("=== 2. 创建指定时间 ===")
	// 创建 2024年1月15日 14点30分0秒
	t := time.Date(2024, 1, 15, 14, 30, 0, 0, time.Local)
	fmt.Println("指定时间:", t.Format("2006-01-02 15:04:05"))
	fmt.Println()

	fmt.Println("=== 3. 字符串转时间（解析）===")
	timeStr := "2024-01-15 14:30:45"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("解析失败:", err)
	} else {
		fmt.Printf("字符串 '%s' 解析为: %v\n", timeStr, parsedTime)
	}
	fmt.Println()

	fmt.Println("=== 4. 时间转字符串（格式化）===")
	now = time.Now()
	fmt.Println("年月日:", now.Format("2006-01-02"))
	fmt.Println("时分秒:", now.Format("15:04:05"))
	fmt.Println("完整格式:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("中文格式:", now.Format("2006年01月02日 15时04分05秒"))
	fmt.Println("星期:", now.Format("2006-01-02 Monday"))
	fmt.Println()

	fmt.Println("=== 5. 时间加减 ===")
	now = time.Now()
	fmt.Println("当前时间:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("1小时后:", now.Add(1*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Println("30分钟前:", now.Add(-30*time.Minute).Format("2006-01-02 15:04:05"))
	fmt.Println("明天:", now.AddDate(0, 0, 1).Format("2006-01-02"))
	fmt.Println("下个月:", now.AddDate(0, 1, 0).Format("2006-01-02"))
	fmt.Println("去年:", now.AddDate(-1, 0, 0).Format("2006-01-02"))
	fmt.Println()

	fmt.Println("=== 6. 时间比较 ===")
	t1 := time.Date(2024, 1, 15, 14, 30, 0, 0, time.Local)
	t2 := time.Date(2024, 1, 16, 14, 30, 0, 0, time.Local)

	fmt.Println("t1:", t1.Format("2006-01-02"))
	fmt.Println("t2:", t2.Format("2006-01-02"))
	fmt.Println("t1 在 t2 之前?", t1.Before(t2))
	fmt.Println("t1 在 t2 之后?", t1.After(t2))
	fmt.Println("t1 等于 t2?", t1.Equal(t2))
	fmt.Println()

	fmt.Println("=== 7. 计算时间差 ===")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2024, 1, 15, 12, 30, 0, 0, time.Local)

	diff := end.Sub(start)
	fmt.Println("开始时间:", start.Format("2006-01-02"))
	fmt.Println("结束时间:", end.Format("2006-01-02"))
	fmt.Println("相差:", diff)
	fmt.Println("相差小时:", diff.Hours())
	fmt.Println("相差分钟:", diff.Minutes())
	fmt.Println("相差天数:", int(diff.Hours()/24))
	fmt.Println()

	fmt.Println("=== 8. 获取时间的各个部分 ===")
	now = time.Now()
	fmt.Println("年:", now.Year())
	fmt.Println("月:", int(now.Month()))
	fmt.Println("日:", now.Day())
	fmt.Println("时:", now.Hour())
	fmt.Println("分:", now.Minute())
	fmt.Println("秒:", now.Second())
	fmt.Println("星期:", int(now.Weekday()))
	fmt.Println()

	fmt.Println("=== 9. 时间戳转换 ===")
	now = time.Now()

	// 时间 -> 时间戳（秒）
	timestamp := now.Unix()
	fmt.Println("当前时间戳（秒）:", timestamp)

	// 时间戳 -> 时间
	tFromTimestamp := time.Unix(timestamp, 0)
	fmt.Println("时间戳转回时间:", tFromTimestamp.Format("2006-01-02 15:04:05"))
	fmt.Println()

	fmt.Println("=== 10. 实用工具函数示例 ===")
	// 判断是否是今天
	fmt.Println("是否是今天:", isToday(now))

	// 获取今天的开始时间
	fmt.Println("今天开始时间:", getDayStart(now).Format("2006-01-02 15:04:05"))

	// 获取今天的结束时间
	fmt.Println("今天结束时间:", getDayEnd(now).Format("2006-01-02 15:04:05"))

	fmt.Println("String方法是什么", time.Now().GoString())
}

// 判断是否是今天
func isToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() &&
		t.Month() == now.Month() &&
		t.Day() == now.Day()
}

// 获取某一天的开始时间（00:00:00）
func getDayStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 获取某一天的结束时间（23:59:59）
func getDayEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
