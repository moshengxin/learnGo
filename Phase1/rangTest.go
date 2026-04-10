package main

import "fmt"

func main() {

	// rang数组
	arr := [3]int{1, 3, 5}
	sum := 0
	for i := range arr {
		fmt.Println(i)
		sum += arr[i]
	}
	fmt.Println(sum)

	// rang map
	mp := map[string]int{"1": 1, "2": 2}
	for k, v := range mp {
		fmt.Println(k, v)
	}

}
