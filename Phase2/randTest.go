package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. 生成两个 0~99 的随机整数（默认种子，每次运行结果可能相同）
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// 2. 生成 0.0 ~ 1.0 之间的随机浮点数
	fmt.Println(rand.Float64())

	// 3. 生成 5.0 ~ 10.0 之间的随机浮点数
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 4. 使用【当前纳秒】作为随机种子，每次运行结果都不同（真实随机）
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 5. 使用固定种子 42，每次运行结果完全一样（固定随机序列）
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print("======", r2.Intn(100), ",")
	fmt.Print("==========", r2.Intn(100))
	fmt.Println()

	// 6. 和上面种子相同，结果完全相同
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
