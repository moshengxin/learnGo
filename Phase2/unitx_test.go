package main

// 单元测试要点：文件名要包含 _test.go，函数要Test开头
/**
t *testing.T是单元测试，函数前缀TestXxx，用于验证函数是否正确，不用手动循环，自动计算耗时
t *testing.B 是基准测试，函数名必须以 Benchmark 开头，参数是 *testing.B，要使用循环测试性能
*/

import (
	"fmt"
	"testing" // Go 内置测试库，专门写单元测试、性能测试
)

// IntMin 返回两个整数中较小的值
// 这是【被测试的业务函数】
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntMinBasic 基础单元测试
// 函数名必须以 Test 开头，参数是 *testing.T
func TestIntMinBasic(t *testing.T) {
	// 调用函数，传入测试用例
	ans := IntMin(2, -2)
	// 判断结果是否符合预期
	if ans != -2 {
		// 测试失败，输出错误信息
		t.Errorf("IntMin(2, -2) = %d; 期望值: -2", ans)
	}
}

// TestIntMinTableDriven 表格驱动测试（Go 推荐写法）
// 一次性测试多组用例，代码更简洁、易扩展
func TestIntMinTableDriven(t *testing.T) {
	// 定义测试用例切片：每组包含 输入a、输入b、期望值
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// 遍历所有测试用例
	for _, tt := range tests {
		// 给每个子测试起名字
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// t.Run 运行子测试，方便单独看每个用例结果
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			// 断言结果是否正确
			if ans != tt.want {
				t.Errorf("实际值: %d, 期望值: %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin 基准测试（性能测试）
// 函数名必须以 Benchmark 开头，参数是 *testing.B
func BenchmarkIntMin(b *testing.B) {
	// b.N 是测试框架自动控制的循环次数，无需手动指定
	for i := 0; i < b.N; i++ {
		// 循环执行被测试函数，统计耗时、性能
		IntMin(1, 2)
	}
}

func TestAdd(t *testing.T) {

}
