package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Rsp struct {
	Code int
	Msg  string
	Time int
	Data string
}

type Tu struct {
	Id         int    `json:"id"`
	Account    string `json:"account"`
	Score      int    `json:"score"`
	Createtime string `json:"createtime"`
}

func main() {

}

func testPost() {

	// 创建http客户端
	client := http.Client{}
	fmt.Println(client)

	// 添加body的json格式参数
	t := new(Tu)
	t.Id = 1
	t.Account = "账户"
	t.Score = 100
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	t.Createtime = format

	jsonData, _ := json.Marshal(t)
	fmt.Println("打印看看", string(jsonData))
	fmt.Println("------------------")
	body := bytes.NewReader(jsonData)

	// 新建post请求
	var rq, _ = http.NewRequest("POST", "http://127.0.0.1:9999/admin/tutelageRecord/test", body)

	// 设置相关请求头
	rq.Header.Add("Content-Type", "application/json")
	fmt.Println(rq)
	fmt.Println("00000000000")
	do, err := client.Do(rq)
	if err != nil {
		fmt.Println("错误不为空，请求出错", err)
		panic(err)
	}
	// 事后延迟关闭流
	closer := do.Body
	defer closer.Close()

	rsp, error := io.ReadAll(do.Body)
	if error != nil {
		fmt.Println("错误不为空，请求出错", error)
		panic(error)
	}
	fmt.Println("响应打印", string(rsp))
}

func testGet() {
	// 创建http客户端
	client := http.Client{}

	// ?添加参数
	map1 := url.Values{}
	map1.Add("phone", "3047012090")
	map1.Add("areaCode", "92")
	parameters := map1.Encode()
	fmt.Println("------", parameters)

	// 创建请求
	request, _ := http.NewRequest("POST", "https://testgetmoney000000.com/money/api/task/encrypt_phone?"+parameters, nil)
	// 添加请求头
	request.Header.Add("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1lIjoiMTgxMDAwMDAwMDAiLCJpZCI6Ijc3ODc2NDk3NTU2NjkxNyIsImV4cCI6MTc3NzYwNDE4MX0.kmv-k2FiQViipYtRdxYqYmZYZqhx0aaWp0JFbEvpg_A")
	// 执行request请求
	rsp, rspErr := client.Do(request)
	if rspErr != nil {
		fmt.Println("错误不为空，请求出错", rspErr)
		panic(rspErr)
	}
	// 延迟自动关闭流
	defer rsp.Body.Close()
	// 读取接口响应内容
	re, _ := io.ReadAll(rsp.Body)
	fmt.Println("响应打印", string(re))

	// json转成结构体
	var rsp1 Rsp
	json.Unmarshal(re, &rsp1)
	fmt.Println("转成结构体后打印看看==", rsp1)
}
