package main

import (
	"fmt"     // 用于打印输出
	"net"     // 用于拆分主机和端口
	"net/url" // Go 标准库，专门解析 URL
)

func main() {
	// 待解析的 URL 字符串（postgres 数据库连接地址）
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析 URL，返回结构化的 URL 对象，若解析失败直接 panic
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 1. 获取 URL 协议（scheme）
	fmt.Println("协议：", u.Scheme)

	// 2. 获取用户认证信息（用户名+密码）
	fmt.Println("完整用户信息：", u.User)
	fmt.Println("用户名：", u.User.Username())
	// 提取密码，忽略错误
	p, _ := u.User.Password()
	fmt.Println("密码：", p)

	// 3. 获取主机 + 端口
	fmt.Println("主机+端口：", u.Host)
	// 拆分主机和端口（net 标准库工具）
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("主机：", host)
	fmt.Println("端口：", port)

	// 4. 获取路径和锚点
	fmt.Println("路径：", u.Path)
	fmt.Println("锚点（#后面）：", u.Fragment)

	// 5. 获取查询参数
	fmt.Println("原始参数：", u.RawQuery)
	// 将原始参数解析为 map 结构
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("解析后的参数map：", m)
	// 取出参数 k 对应的第一个值
	fmt.Println("参数k的值：", m["k"][0])
}
