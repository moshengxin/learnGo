package main

import (
	"fmt"
	"sort"
)

func main() {

	s := make([]string, 3)
	s1 := []string{}
	fmt.Println("emp:", s, s1)

	s[0] = "z"
	s[1] = "m"
	s[2] = "x"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var slsort = make([]int, 3)
	slsort[0] = 5
	slsort[1] = 2
	slsort[2] = 3
	fmt.Println("排序前:", slsort)
	// slsort排序一下
	//for i := 0; i < len(slsort); i++ {
	//	for j := i + 1; j < len(slsort); j++ {
	//		if slsort[i] > slsort[j] {
	//			slsort[i], slsort[j] = slsort[j], slsort[i]
	//		}
	//	}
	//}
	sort.Ints(slsort)
	sort.Sort(sort.IntSlice(slsort))
	fmt.Println("排序后:", slsort)
}
