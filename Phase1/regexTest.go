package main

// 导入需要的包：字节操作、打印、正则表达式
import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 1. 直接匹配：判断字符串是否满足正则 p([a-z]+)ch
	// 正则含义：p开头 + 一个以上小写字母 + ch结尾
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // 输出：true

	// 2. 编译正则（提前编译，效率更高）
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 3. 判断是否匹配字符串 peach
	fmt.Println(r.MatchString("peach")) // 输出：true

	// 4. 查找第一个匹配的字符串
	fmt.Println(r.FindString("peach punch")) // 输出：peach

	// 5. 查找第一个匹配的开始/结束索引 [开始, 结束]
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // 输出：idx: [0 5]

	// 6. 子匹配：返回【完整匹配】+【括号里捕获的内容】
	fmt.Println(r.FindStringSubmatch("peach punch")) // 输出：[peach ea]

	// 7. 子匹配的索引：完整匹配索引 + 子匹配索引
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // 输出：[0 5 1 3]

	// 8. 查找所有匹配项，-1 表示返回全部
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // 输出：[peach punch pinch]

	// 9. 所有子匹配的索引
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 10. 只返回前 2 个匹配项
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // 输出：[peach punch]

	// 11. 也可以匹配字节数组 []byte 类型
	fmt.Println(r.Match([]byte("peach"))) // 输出：true

	// 12. MustCompile：编译正则，出错直接 panic（用于全局变量）
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) // 输出正则表达式

	// 13. 替换所有匹配项：把 peach 替换成 <fruit>
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // 输出：a <fruit>

	// 14. 函数式替换：把匹配到的内容传入函数处理（这里转大写）
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // 匹配内容转大写
	fmt.Println(string(out))                   // 输出：a PEACH
}
