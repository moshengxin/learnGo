package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template" // Go 标准库文本模板引擎，用来动态渲染文本
)

func main() {
	//test11()

	moseyTemplate := template.New("mosey")
	str := "验证码是：{{.}}"
	moseyTemplate.Parse(str)
	var buf bytes.Buffer
	moseyTemplate.Execute(&buf, "456")
	fmt.Println("====", buf.String())

}

func test11() {
	// 1. 创建一个名为 t1 的模板对象
	t1 := template.New("t1")

	// 解析模板字符串
	// {{.}} 代表：后面 Execute 传进来的【整个数据】
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil { // 解析失败直接panic
		panic(err)
	}

	// 重点：template.Must
	// 作用：自动接管 err，解析出错直接 panic，不用自己写if判断err
	// 覆盖掉上面旧的 t1，模板内容更新为新的
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板，输出到控制台 os.Stdout
	// 传入字符串
	t1.Execute(os.Stdout, "some text")
	// 传入数字
	t1.Execute(os.Stdout, 5)
	// 传入字符串切片
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 封装工具函数：快速创建模板，不用重复写 New+Parse+Must
	// name:模板名  t:模板内容
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 2. 模板取值：结构体 / map 的字段 {{.Name}}
	t2 := Create("t2", "Name: {{.Name}}\n")

	// 传入匿名结构体，模板自动取 .Name 字段
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 传入map，模板自动取 key 为 Name 的值
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// 3. 模板内 if 判断语法 {{if}} {{else}} {{end}}
	// {{- -}} 作用：去除模板自带的多余空格、换行，让输出更干净
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	// 传入非空字符串 → 条件成立，输出 yes
	t3.Execute(os.Stdout, "not empty")
	// 传入空字符串 → 条件不成立，输出 no
	t3.Execute(os.Stdout, "")

	// 4. 模板内循环遍历 {{range}} {{end}}
	// range 专门遍历切片、数组
	t4 := Create("t4",
		"Range: {{range .}}{{.}} == {{end}}\n")

	// 传入字符串切片，range会逐个遍历元素输出
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
