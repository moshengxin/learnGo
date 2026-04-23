package main

// 导入包：
// encoding/json：Go 官方的 JSON 处理库（序列化、反序列化）
// fmt：打印输出
// os：系统输入输出，这里用于输出到控制台
import (
	"encoding/json"
	"fmt"
	"os"
)

// 定义结构体 response1
// 没有写 json 标签，序列化时字段名会**原样输出**
type response1 struct {
	Page   int      // 字段名：Page
	Fruits []string // 字段名：Fruits
}

// 定义结构体 response2
// 带有 `json:"page"` 标签：序列化时字段名会变成**小写 page、fruits**
type response2 struct {
	Page   int      `json:"page"`   // 序列化后键名：page
	Fruits []string `json:"fruits"` // 序列化后键名：fruits
}

func main() {

	// --------------------------
	// 1. json.Marshal：把 Go 类型 → JSON 字符串（序列化）
	// --------------------------

	// 布尔值 → JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) // 输出：true

	// 整数 → JSON
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 输出：1

	// 浮点数 → JSON
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 输出：2.34

	// 字符串 → JSON
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // 输出："gopher"

	// 切片（数组）→ JSON 数组
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // 输出：["apple","peach","pear"]

	// map → JSON 对象
	mapA := map[string]string{}
	mapA["a"] = "1"
	mapA["b"] = "2"

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) // 输出：{"apple":5,"lettuce":7}

	// 结构体 response1 → JSON（字段名大写），  &代表指针类型
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // 输出：{"Page":1,"Fruits":["apple","peach","pear"]}

	// 结构体 response2 → JSON（带标签，字段名小写）  &代表指针类型
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // 输出：{"page":1,"fruits":["apple","peach","pear"]}

	// --------------------------
	// 2. json.Unmarshal：JSON 字符串 → Go 类型（反序列化）
	// --------------------------

	// 准备一段 JSON 数据（byte 数组）
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 定义一个 map，用来接收解析后的数据
	// value 是 interface{}：可以接收任意类型
	var dat map[string]interface{}

	// 解析 JSON 到 dat 变量中
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("我在dat", dat) // 输出解析后的 map

	// 类型断言：把解析出来的 num 转成 float64
	num := dat["num"].(float64)
	fmt.Println(num) // 输出：6.13

	// 类型断言：把 strs 转成 []interface{} 切片
	strs := dat["strs"].([]interface{})
	// 再把切片里的第一个元素转成 string
	str1 := strs[0].(string)
	fmt.Println(str1) // 输出：a

	// --------------------------
	// 3. JSON 直接解析到结构体（最常用！）
	// --------------------------
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	// 直接解析到结构体 res
	json.Unmarshal([]byte(str), &res) // 这里加&指针，才能给res赋值
	fmt.Println("我是res", res)         // 输出结构体完整内容
	fmt.Println(res.Fruits[0])        // 输出：apple

	// --------------------------
	// 4. json.NewEncoder：直接把数据编码成 JSON 并输出到控制台
	// --------------------------
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // 直接打印输出：{"apple":5,"lettuce":7}
}
